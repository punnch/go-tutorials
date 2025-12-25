package main

import (
	"fmt"
	"time"
)

// Task:

// 1. Create 3 chan (int, str, float64)

// 2. Create 3 goroutines
// 1) Every goroutine recives a chan -> writes text in 'for loop'
// 2) Interval:
// - int 300ms
// - int 1s
// - float64 5s

// 4. Read a chan value -> print

func main() {
	intChan := make(chan int)
	strChan := make(chan string)
	floatChan := make(chan float64)

	// Int chan
	go func() {
		for {
			intChan <- 1
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// Str chan
	go func() {
		for {
			strChan <- "I'm punnch"
			time.Sleep(1 * time.Second)
		}
	}()

	// Float chan
	go func() {
		for {
			floatChan <- 3.333
			time.Sleep(5 * time.Second)
		}

	}()

	for {
		select {
		case ch1 := <-intChan:
			fmt.Println(ch1)
		case ch2 := <-strChan:
			fmt.Println(ch2)
		case ch3 := <-floatChan:
			fmt.Println(ch3)
		}
	}
}
