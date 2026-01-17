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

func idHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	msg, ok := noteops.GetNoteById(idNote)
	if ok {
		b, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(b); err != nil {
			fmt.Println("fail http server error:", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)

		msg := "fail to show, no message with provided id"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	if len(notes.Messages) == 0 {
		msg := "no messages yet"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	b, err := json.Marshal(notes.Messages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(b)); err != nil {
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

	msg := "message added"
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Println("fail http server error:", err)
	}
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

		msg := "fail to convert http request body to int:" + err.Error()
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}

	mtx.Lock()
	defer mtx.Unlock()

	ok := noteops.DeleteNoteById(idNote)
	if ok {
		msg := "message deleted"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)

		msg := "fail to show, no message with provided id"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
}

func main() {
	http.HandleFunc("/id", idHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/delete", deleteHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail server error:", err)
	}
}
