package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2) // sleeps 2 seconds, main sleeps 5 seconds
	fmt.Println("***sending value***")
	myChannel <- "a" // returned a value and waits before main recieve
	fmt.Println("***sending value***")
	myChannel <- "b" // executes while main gets that value
}

func main() {
	// random goroutine execution
	myChannel := make(chan string)
	go send(myChannel)

	reportNap("receiving goroutine", 5) // sleeps 5 senconds -> send goroutine will wait

	// main gets values from channels
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
