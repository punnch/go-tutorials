package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	m   map[string]string
	mtx sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]string),
	}
}

func SetKey(wg *sync.WaitGroup, s *SafeMap, k, v string) {
	defer wg.Done()

	s.mtx.Lock()
	s.m[k] = v
	s.mtx.Unlock()
}

func GetKey(wg *sync.WaitGroup, s *SafeMap, k string) {
	defer wg.Done()

	s.mtx.RLock()
	_ = s.m[k]
	s.mtx.RUnlock()
}

func main() {
	wg := &sync.WaitGroup{}

	storage := NewSafeMap()

	keys := []string{"punnch", "nilchan", "german", "cumpot", "kopatich"}
	values := []string{"cool", "loved", "kind", "the best", "bear"}

	goroutineNum := len(keys)

	initTime := time.Now()

	for i := range goroutineNum {
		wg.Add(1)
		go SetKey(wg, storage, keys[i], values[i])
	}

	for i := range goroutineNum {
		wg.Add(1)
		go GetKey(wg, storage, keys[i])
	}

	wg.Wait()

	fmt.Println("Time:", time.Since(initTime))

	storage.mtx.RLock()
	fmt.Println(len(storage.m))
	fmt.Println(storage.m)
	storage.mtx.RUnlock()
}
