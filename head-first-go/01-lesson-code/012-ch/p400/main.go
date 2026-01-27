package main

import "fmt"

func main() {
	defer fmt.Println("Deffered in one")
	defer fmt.Println("Deffered in two")
	panic("Something")
}
