package main

import (
	"fmt"
	"net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	num := r.URL.Query().Get("num")
	str := r.URL.Query().Get("str")

	fmt.Println("num:", num)
	fmt.Println("str:", str)
}

func main() {
	http.HandleFunc("/query", queryHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("fail http server error:", err)
	}
}
