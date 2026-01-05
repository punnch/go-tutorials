package main

import (
	"fmt"
	"sync"
)

var slice []int
var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1_000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Before:", len(slice))

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go increase(wg)
	}

	wg.Wait()

	// Expected: number = 10_000
	fmt.Println("After:", len(slice))
}
