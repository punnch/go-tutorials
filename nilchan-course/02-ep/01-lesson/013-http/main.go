package main

import (
	"fmt"
	"net/http"
	"time"
)

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	str := "Sleep operation was successful!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Sleep err:", err)
	} else {
		fmt.Println("Sleep HTTP request was successful!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Payment was canceled successful!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Payment cancel err:", err)
	} else {
		fmt.Println("Payment cancel HTTP request was successful!")
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Payment was successful"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Payment err", err)
	} else {
		fmt.Println("Payment HTTP request was successful!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("'Hello world' err", err)
	} else {
		fmt.Println("Greeting HTTP request was successful")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)
	http.HandleFunc("/sleep", sleepHandler)

	fmt.Println("Server is started!")

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Error during server work:", err)
	}
}
