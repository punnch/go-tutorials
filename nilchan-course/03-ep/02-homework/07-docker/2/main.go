package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, "Hello from docker!")
		time.Sleep(time.Second)
	}
}
