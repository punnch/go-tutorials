package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	money          int = 1000
	paymentHistory []Payment
	mtx            sync.Mutex
)

type HttpResponse struct {
	Money          int       `json:"money"`
	PaymentHistory []Payment `json:"pHistory"`
}

type Payment struct {
	Description string    `json:"description"`
	USD         int       `json:"usd"`
	FullName    string    `json:"fullName"`
	Address     string    `json:"address"`
	Created     time.Time `json:"created"`
	Canceled    bool      `json:"isCanceled"`
}

func (p Payment) Validate() bool {
	switch {
	case p.Description == "":
		return false
	case p.USD == 0:
		return false
	case p.FullName == "":
		return false
	case p.Address == "":
		return false
	default:
		return true
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error:", err)
		return
	}

	payment.Created = time.Now()

	mtx.Lock()
	defer mtx.Unlock()

	if !payment.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("error:", "validation failed")
		return
	}

	if money-payment.USD >= 0 {
		money -= payment.USD
	} else {
		payment.Canceled = true
	}

	paymentHistory = append(paymentHistory, payment)

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error:", err)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
