package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, name string) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон и отнес газету", name, "в", i, "раз")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg := sync.WaitGroup{}

	// 1
	wg.Add(1)
	go postman(&wg, "'Новости'")

	// 2
	wg.Add(1)
	go postman(&wg, "'Бравл старс'")

	// 3
	wg.Add(1)
	go postman(&wg, "'Go'")

	wg.Wait()

	fmt.Println("main завершился")
}
