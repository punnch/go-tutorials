package main

import (
	"fmt"
	"sync"
)

var email []string

func sendEmail(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		email = append(email, "hello, Punnch! I'm your fan!")
	}
}

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("Before:", len(email))

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go sendEmail(wg)
	}
	wg.Wait()

	// Expected: 10000
	fmt.Println("After:", len(email))

	// Conclusion: race condition!
}
