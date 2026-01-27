package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var (
	money int = 1000
	bank  int
	mtx   sync.Mutex
)

func payHandler(w http.ResponseWriter, r *http.Request) {
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

	paymentRequest, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert request body to int:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	mtx.Lock()
	if money >= paymentRequest {
		money -= paymentRequest

		msg := "Successful payment:" + strconv.Itoa(paymentRequest) + " Money:" + strconv.Itoa(money)
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to pay with insufficient money"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
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

	saveRequst, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert request body to int:" + err.Error()
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
		return
	}

	mtx.Lock()
	if money >= saveRequst {
		money -= saveRequst
		bank += saveRequst

		msg := "Successful bank depositing:" + strconv.Itoa(saveRequst) + " Bank money:" + strconv.Itoa(bank) + " Money:" + strconv.Itoa(money)
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to deposit with insufficient money"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail http server error:", err)
		}
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Http server error:", err)
	}
}
