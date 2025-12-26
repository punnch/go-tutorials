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
	reportNap("sending goroutine", 2) // спит 2 сек, a main 5 сек -> вкл. быстрее
	fmt.Println("***sending value***")
	myChannel <- "a" // отправил значение и ждет, пока main получит
	fmt.Println("***sending value***")
	myChannel <- "b" // пока main не получит a, эта команда не выполнится
}

func main() {
	// планировщик бахает на рандом какую горутину запускать
	myChannel := make(chan string)
	go send(myChannel)
	// main - получающая горутина, выполняет команду, пока не выполнит - не получит
	reportNap("receiving goroutine", 5) // спит 5 сек -> сенд горутина будет ждать

	// main получает значение канала
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
