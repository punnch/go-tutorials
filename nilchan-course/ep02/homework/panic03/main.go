package main

import (
	"fmt"
	"log"
)

// Task:

// 1. Divide by zero
// 2. Write a value in a nil-map
// 3. Call n+1 el in slice (len is n)
// 4. Call panic manually

// - watch console output
// - try to cath a panic

func main() {
	defer func() {
		p := recover()
		if p != nil {
			log.Fatal(p)
		}
	}()

	// 1
	zero := 0
	division := 1 / zero

	fmt.Println(division)

	// 2
	m := make(map[int]int)
	m = nil
	m[1] = 2

	// // 3
	slice := []int{1, 2, 3}
	fmt.Println(slice[9])

	// 4
	panic("panic called manually!")
}
