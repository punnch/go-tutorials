package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int

	myInt = 42
	myIntPointer = &myInt

	fmt.Println(*myIntPointer)

}
