package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gardener(wg *sync.WaitGroup, area string) {
	defer wg.Done()

	fmt.Println("Я огородник", rand.Intn(5)+1, "и удобряю", area)
	time.Sleep(time.Duration(500+rand.Intn(501)) * time.Millisecond)
}

func main() {
	wg := &sync.WaitGroup{}

	// 1
	wg.Add(1)
	go gardener(wg, "лилии")

	// 2
	wg.Add(1)
	go gardener(wg, "розы")

	// 3
	wg.Add(1)
	go gardener(wg, "ромашки")

	wg.Wait()

	fmt.Println("Огородники закончили свою работу!")
}
