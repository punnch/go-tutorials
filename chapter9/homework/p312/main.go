package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "+", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "-", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	// 1 value
	nine := Number(9)
	nine.Add(7)
	nine.Subtract(3)

	// 2 value
	zero := Number(0)
	zero.Add(0)
	zero.Subtract(3)
}
