// Task:

// 1. Create a func that generated a random num
// 1) Recieves:
// - chan int
// 2) Writes a num in a recieved channel

// 2. Start 3  goroutines

// 3. Read values that chan recieves

// 4. Print recieved values

// 5. Make sure:
// - goroutines passed the values
// - values are read

package main

import (
	"fmt"
	"math/rand"
)

func randomNum(ch chan int) {
	num := rand.Int()
	ch <- num
}

func main() {
	ch := make(chan int)
	slice := []int{}
	var val int

	for range 3 {
		go randomNum(ch)
	}

	for range 3 {
		val = <-ch
		slice = append(slice, val)
	}

	for i, v := range slice {
		fmt.Printf("%d - %d\n", i+1, v)
	}
}
