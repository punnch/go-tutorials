package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	messages []string
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
	messages = append(messages, message)
	mtx.Unlock()

	fmt.Println(message)
}

func main() {
	http.HandleFunc("/message", messageHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail server error:", err)
	}
}
