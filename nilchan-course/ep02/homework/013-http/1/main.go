package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Punnch is the best guy!"))
	if err != nil {
		fmt.Println("Request err:", err)
	} else {
		fmt.Println("Requset was successful!")
	}
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error during server work")
	}
}
