package main

import "fmt"

func main() {
	var subscriber struct {
		name   string
		rating float64
		active bool
	}

	subscriber.name = "Dima"
	subscriber.rating = 10.0
	subscriber.active = true

	fmt.Println("Name:", subscriber.name)
	fmt.Println("Rating:", subscriber.rating)
	fmt.Println("Active?", subscriber.active)
}
