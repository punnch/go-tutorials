package main

import (
	"fmt"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("fail method not allowed")
		return
	}

	msg := "true request! (get)"
	fmt.Println(msg)
}

func main() {
	http.HandleFunc("/method", methodHandler)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
