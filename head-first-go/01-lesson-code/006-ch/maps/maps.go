package main

import "fmt"

func main() {
	// 1 example (make)
	ranks := make(map[string]int)

	ranks["silver"] = 1
	ranks["gold"] = 2
	ranks["diamond"] = 3

	fmt.Println(ranks["silver"])
	fmt.Println(ranks["diamond"])

	// 2 example of making map without "make"
	myMap := map[string]float64{"some": 1.0, "fucking": 3.0}

	fmt.Println("MyMap: ", myMap)

	// 3 example
	numbers := make(map[string]float64)
	numbers["assigned"] = 16.41341

	fmt.Printf("%.2f \n", numbers["assigned"])
	fmt.Println(numbers["some"])
}
