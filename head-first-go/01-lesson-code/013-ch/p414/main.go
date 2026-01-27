package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	URL  string
	Size int
}

func main() {
	results := make(chan Result)
	urls := []string{"https://go.dev", "https://chatgpt.com", "https://github.com"}

	for _, url := range urls {
		go responseSize(url, results)
	}

	for range urls {
		res := <-results
		fmt.Printf("%s - %d bytes\n", res.URL, res.Size)
	}
}

// responseSize uses net/http and gets url to return byte size of page
func responseSize(url string, ch chan Result) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	ch <- Result{URL: url, Size: len(body)}
}
