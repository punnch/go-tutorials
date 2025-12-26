package main

import "fmt"

func main() {
	var pet struct {
		name string
		age  int
	}

	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)
}
