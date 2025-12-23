package main

import (
	"fmt"
	"time"
)

// Func mine() describes miner's trip to the mine
// 1. Recieves:
// - number of a trip
// 2. Returns:
// - mined coal
func mine(transferPoint chan int, n int) {
	fmt.Println("A mine trip number", n, "started...")
	time.Sleep(1 * time.Second)
	fmt.Println("A mine trip number", n, "ended!")

	transferPoint <- 10
}

func main() {
	coal := 0
	// Make channel
	transferPoint := make(chan int)

	// Initial time before a test
	initTime := time.Now()

	// Init 3 goroutines
	for i := 1; i <= 3; i++ {
		go mine(transferPoint, i)
	}

	// Add a value from chan to a variable
	for i := 1; i <= 3; i++ {
		coal += <-transferPoint
	}

	fmt.Println("Mined coal amount:", coal)

	// Test of the time spent on the function execution
	fmt.Println("Mining time:", time.Since(initTime))
}
