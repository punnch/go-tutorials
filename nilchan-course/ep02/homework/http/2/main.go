package main

import (
	"fmt"
	"net/http"
)

func dogHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("I'm a dog and say 'Ruff'"))
	if err != nil {
		fmt.Println("Dog HTTP request error:", err)
	} else {
		fmt.Println("Dog HTTP request was successful!")
	}
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("I'm a cat and say 'Meow'"))
	if err != nil {
		fmt.Println("Cat HTTP request error:", err)
	} else {
		fmt.Println("Cat HTTP request was successful!")
	}
}

func cowHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("I'm a cow and say 'Moo'"))
	if err != nil {
		fmt.Println("Cow HTTP request error:", err)
	} else {
		fmt.Println("Cow HTTP request was successful!")
	}
}

func main() {
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/cow", cowHandler)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error during server work")
	}
}
