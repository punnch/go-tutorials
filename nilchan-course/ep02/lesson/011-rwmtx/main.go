package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0
var mtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()

	// Expected: 1_000_000
	fmt.Println("Likes:", likes)
	fmt.Println(time.Since(initTime))
}
