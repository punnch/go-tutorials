package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSleep() int {
	randomNum := 1 + rand.Intn(10)
	miliseconds := randomNum * 100
	time.Sleep(time.Duration(miliseconds) * time.Millisecond)

	return miliseconds
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 1 goroutine
	go func() {
		sleepTime := randomSleep()
		fmt.Println("Goroutine 1 slept:", sleepTime, "ms")
		ch1 <- sleepTime
	}()

	// 2 goroutine
	go func() {
		sleepTime := randomSleep()
		fmt.Println("Goroutine 2 slept:", sleepTime, "ms")
		ch2 <- sleepTime
	}()

	time.Sleep(500 * time.Millisecond)

	select {
	case value1 := <-ch1:
		fmt.Println("Goroutine 1 get value:", value1)
	case value2 := <-ch2:
		fmt.Println("Goroutine 2 get value:", value2)
	default:
		fmt.Println("No value")
	}
}
