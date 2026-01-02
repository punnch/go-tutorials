package main

import (
	"fmt"
	"time"
)

// Task:

// Init 3 goroutins with random text

func main() {
	go func() {
		for {
			fmt.Println("My life")
		}
	}()

	go func() {
		for {
			fmt.Println("be")
		}
	}()

	go func() {
		for {
			fmt.Println("like...")
		}
	}()

	time.Sleep(10 * time.Millisecond)
}
