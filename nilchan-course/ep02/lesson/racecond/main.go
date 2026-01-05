package main

import (
	"fmt"
	"sync"
)

var number int

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		number++
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Before:", number)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go increase(wg)
	}

	wg.Wait()

	// Expected: number = 5_000
	fmt.Println("After:", number)
}
