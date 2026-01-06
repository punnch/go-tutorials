package main

import (
	"fmt"
	"sync"
)

var candidate int

func ClassVotes(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 30; i++ {
		candidate++
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Before:", candidate)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go ClassVotes(wg)
	}
	wg.Wait()

	// Expected: 3000
	fmt.Println("After:", candidate)

	// Conclusion: race condition!
}
