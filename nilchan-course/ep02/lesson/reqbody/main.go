package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	money int = 50
	bank  int
	mtx   sync.Mutex
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail http request body error:", err)
		return
	}

	bodyStr := string(httpRequestBody)

	paymentRequest, err := strconv.Atoi(bodyStr)
	if err != nil {
		fmt.Println("fail to convert request body to int:", err)
		return
	}

	mtx.Lock()
	if money >= paymentRequest {

		time.Sleep(3 * time.Second)

		money -= paymentRequest

		fmt.Println("Успешное оплата:", paymentRequest)
		fmt.Println("Денег на счету:", money)
	} else {
		fmt.Println("Недостаточно средств для оплаты!")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail http request body error:", err)
		return
	}

	bodyStr := string(httpRequestBody)

	saveRequst, err := strconv.Atoi(bodyStr)
	if err != nil {
		fmt.Println("fail to convert request body to int:", err)
		return
	}

	mtx.Lock()
	if money >= saveRequst {
		money -= saveRequst
		bank += saveRequst

		fmt.Println("Успешное зачисление в банк:", saveRequst)
		fmt.Println("Денег в банке:", bank)
		fmt.Println("Денег на счету:", money)
	} else {
		fmt.Println("Недостаточно средств для перевода в банк!")
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
