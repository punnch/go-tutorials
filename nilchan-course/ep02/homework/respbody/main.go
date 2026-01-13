package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func responseHeader(w http.ResponseWriter, r *http.Request) {
	httpResuestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail http server error:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	number, err := strconv.Atoi(string(httpResuestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert htm request body to int:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	switch number {
	case 2:
		msg := "2 = return status code 200"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}

	case 3:
		w.WriteHeader(http.StatusMultipleChoices)

		msg := "3 = return status code 300"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	case 4:
		w.WriteHeader(http.StatusBadRequest)

		msg := "4 = return status code 400"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
}

func main() {
	http.HandleFunc("/response", responseHeader)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
