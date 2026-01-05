package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gardener(wg *sync.WaitGroup, n int) {
	defer wg.Done()

	fmt.Println("Я огородник", n, "начинаю полив...")
	time.Sleep(time.Duration(500+rand.Intn(501)) * time.Millisecond)
	fmt.Println("Я огородник", n, "полив закончил!")
}

func main() {
	min := 3
	max := 10

	// range: 3-10
	gardeners := min + rand.Intn(max-min+1)

	wg := &sync.WaitGroup{}

	for i := 1; i <= gardeners; i++ {
		wg.Add(1)
		go gardener(wg, i)
	}

	wg.Wait()

	fmt.Println("-----------------------------")
	fmt.Println("Огородники закончили свою работу!")
}
