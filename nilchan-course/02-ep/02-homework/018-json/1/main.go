package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bio struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Age       int     `json:"age"`
	IsMarried bool    `json:"isMarried"`
	Height    float64 `json:"height"`
}

func (b Bio) Validate() bool {
	switch {
	case b.Name == "":
		return false
	case b.Address == "":
		return false
	case b.Age == 0:
		return false
	case b.Height == 0.0:
		return false
	default:
		return true
	}
}

func (b Bio) Println() {
	fmt.Println("Name:", b.Name)
	fmt.Println("Address:", b.Address)
	fmt.Println("Age:", b.Age)
	fmt.Println("IsMarried?:", b.IsMarried)
	fmt.Println("Height:", b.Height)
}

func bioHandler(w http.ResponseWriter, r *http.Request) {
	var bio Bio
	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !bio.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bio.Println()
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	bio := Bio{
		Name:      "YourName",
		Address:   "Russia, Moscow, Red Square 10",
		Age:       18,
		IsMarried: false,
		Height:    180,
	}

	b, err := json.Marshal(bio)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("fail http server error:", err)
	}
}

func main() {
	http.HandleFunc("/bio", bioHandler)
	http.HandleFunc("/template", templateHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
