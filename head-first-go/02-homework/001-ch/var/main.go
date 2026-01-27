package main

import "fmt"

func main() {
	// 1 task
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")

	// 2 task
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08

	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")

	fmt.Println("Within budget?", total <= float64(availableFunds))
}
