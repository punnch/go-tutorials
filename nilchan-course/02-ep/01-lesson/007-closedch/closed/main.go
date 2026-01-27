package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transferPoint := make(chan int)
	coal := 0

	go func() {
		// range: 3-10
		iterations := 3 + rand.Intn(8)
		fmt.Println("Mine trips:", iterations)

		for i := 1; i <= iterations; i++ {
			time.Sleep(1 * time.Second)

			// + 10 coal for each trip
			transferPoint <- 10
		}

		close(transferPoint)
	}()

	for v := range transferPoint {
		coal += v
		fmt.Println("Coal:", coal)
	}

	fmt.Println("Miner is done!")
	fmt.Println("-----------------")
	fmt.Println("Total coal:", coal)
}
