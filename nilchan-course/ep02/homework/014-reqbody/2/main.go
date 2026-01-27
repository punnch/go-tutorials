package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/k0kubun/pp"
)

var (
	messages map[int]string
	id       int
	mtx      sync.Mutex
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail http request body error:", err)
		return
	}

	message := string(httpRequestBody)

	mtx.Lock()
	id++
	messages[id] = message
	mtx.Unlock()

	fmt.Println("Add:", id, "-", message)

	mtx.Lock()
	pp.Println("Messages:", messages)
	mtx.Unlock()
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail http request body error:", err)
		return
	}

	str := string(httpRequestBody)

	idMessage, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("fail to convert http request body to int:", err)
		return
	}

	mtx.Lock()
	_, ok := messages[idMessage]
	if ok {
		fmt.Println("Delete:", idMessage, "-", messages[idMessage])
		delete(messages, idMessage)
		pp.Println("Messages:", messages)
	} else {
		fmt.Println("fail to delete, no message with provided id")
	}
	mtx.Unlock()
}

func main() {
	messages = make(map[int]string)

	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/delete", deleteHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail server error:", err)
	}
}
