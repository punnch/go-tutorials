package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, name string) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("I'm postman and carry a magazine", name, "in", i, "time")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	// 1
	wg.Add(1)
	go postman(wg, "'News'")

	// 2
	wg.Add(1)
	go postman(wg, "'Brawls stars'")

	// 3
	wg.Add(1)
	go postman(wg, "'Go'")

	wg.Wait()

	fmt.Println("main ended")
}
