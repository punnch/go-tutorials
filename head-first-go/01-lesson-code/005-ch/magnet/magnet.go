package main

import "fmt"

func sum(numbers ...int) int {
	var sum int = 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(sum(4, 7, 1))
	fmt.Println(sum(9, 2))
}
