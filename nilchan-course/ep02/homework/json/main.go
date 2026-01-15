package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	money          int = 1000
	paymentHistory []Payment
	mtx            sync.Mutex
)

type Payment struct {
	Description string
	USD         int
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	parts := strings.SplitN(string(httpRequestBody), ",", 2)

	usd, err := strconv.Atoi(parts[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payment := Payment{
		Description: parts[1],
		USD:         usd,
	}

	mtx.Lock()
	if money-usd >= 0 {
		money -= usd
	}

	paymentHistory = append(paymentHistory, payment)

	fmt.Println("Money:", money)
	fmt.Println("History:", paymentHistory)
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
