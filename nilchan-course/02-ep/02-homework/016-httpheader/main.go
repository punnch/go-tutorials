package main

import (
	"fmt"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("k:", k, " -- ", "v:", v)
	}

	name := r.Header.Get("MyName")
	fmt.Println("Hello,", name, "!")
}

func main() {
	http.HandleFunc("/header", headerHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
