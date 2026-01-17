package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"study/noteops"
	"study/notes"
)

var mtx sync.Mutex

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	result := notes.Messages

	if len(result) == 0 {
		msg := "no messages yet"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	// Querys
	q := r.URL.Query()

	idStr := q.Get("id")
	urgentStr := q.Get("urgent")

	// Id query (int)
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		note, ok := noteops.GetNoteById(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// array with one note
		result = []notes.Note{note}
	}

	// Urgent query (bool)
	if urgentStr != "" {
		urgent, err := strconv.ParseBool(urgentStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var filtered = make([]notes.Note, 0)

		for _, v := range result {
			if v.Urgent == urgent {
				filtered = append(filtered, v)
			}
		}
		result = filtered
	}

	// Send messages
	b, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("fail http server error:", err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var note notes.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	note.Created = time.Now()

	if !noteops.Validate(note) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mtx.Lock()
	notes.Id++
	note.ID = notes.Id
	notes.Messages = append(notes.Messages, note)
	mtx.Unlock()
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	idNote, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	ok := noteops.DeleteNoteById(idNote)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/delete", deleteHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail server error:", err)
	}
}
