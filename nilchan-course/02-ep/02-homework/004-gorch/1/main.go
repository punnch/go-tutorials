// Task:

// 1. Create a func
// 1) Recieves:
// - own number int
// 2) Prints 5 times with 1s interval:
// "I'm goroutine N and output to the screen X time"

// 2. Start 3 goroutines

package main

import (
	"fmt"
	"time"
)

func TestFunc(n int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("I'm goroutine", n, "and output to the screen", i, "time")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go TestFunc(1)
	go TestFunc(2)
	go TestFunc(3)

	time.Sleep(6 * time.Second)
}
