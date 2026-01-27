package main

import "fmt"

func main() {
	// nil chan
	var ch chan int

	go func() {
		ch <- 1
	}()

	v := <-ch
	fmt.Println(v)
}

// conclusion: don't use nil channels
