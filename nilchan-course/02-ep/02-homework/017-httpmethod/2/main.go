package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var (
	messages map[int]string
	id       int
	mtx      sync.Mutex
)

func idHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "fail not allowed method"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)

		}
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail http server error:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}

	idMessage, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert http request body to int:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	mtx.Lock()
	_, ok := messages[idMessage]
	if ok {
		msg := messages[idMessage]
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)

		msg := "fail to show, no message with provided id"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
	mtx.Unlock()
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "fail not allowed method"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)

		}
		return
	}

	var str string

	mtx.Lock()
	for k, v := range messages {
		str += fmt.Sprintf("\n%d - %s", k, v)
	}
	mtx.Unlock()

	msg := "Messages:" + str
	fmt.Println(msg)

	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Println("fail http server error:", err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "fail not allowed method"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)

		}
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail http server error:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	message := string(httpRequestBody)

	mtx.Lock()
	id++
	messages[id] = message
	mtx.Unlock()

	msg := "Add:" + strconv.Itoa(id) + " - " + message
	fmt.Println(msg)

	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Println("fail http server error:", err)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "fail not allowed method"
		fmt.Println(msg)
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)

		}
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail http server error:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	idMessage, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert http request body to int:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	mtx.Lock()
	_, ok := messages[idMessage]
	if ok {
		msg := "Delete:" + strconv.Itoa(idMessage) + " - " + messages[idMessage]
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		delete(messages, idMessage)
	} else {
		w.WriteHeader(http.StatusNotFound)

		msg := "fail to delete, no message with provided id"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
	mtx.Unlock()
}

func main() {
	messages = make(map[int]string)

	http.HandleFunc("/id", idHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/delete", deleteHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail server error:", err)
	}
}
