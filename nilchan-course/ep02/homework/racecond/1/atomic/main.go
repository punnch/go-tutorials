package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var candidate atomic.Int64

func ClassVotes(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 30; i++ {
		candidate.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Before:", candidate.Load())

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go ClassVotes(wg)
	}
	wg.Wait()

	// Expected: 3000
	fmt.Println("After:", candidate.Load())
}
