package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

type parts struct {
	description string
	count       int
}

func main() {
	// Simple declaration
	porshe := car{}
	porshe.name = "Porshe"
	porshe.topSpeed = 248
	fmt.Println("Name:", porshe.name)
	fmt.Println("Speed:", porshe.topSpeed)

	// Var declaration
	var bolts parts
	bolts.description = "Porshe parts"
	bolts.count = 16
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

}
