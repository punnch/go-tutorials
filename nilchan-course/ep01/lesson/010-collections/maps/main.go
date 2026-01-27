package main

import "fmt"

func main() {
	// 1
	shop := map[string]float64{
		"bread":     0.99,
		"juice":     1,
		"ice_cream": 3,
		"chocolate": 5,
	}

	fmt.Println(shop)
	fmt.Println("cost:", shop["chocolate"], "$")

	shop["cake"] = 10

	fmt.Println(shop)
	fmt.Println("cost:", shop["cake"], "$")

	// Call non-created key
	fmt.Println(shop["not_created"])

	// Check output
	burg, ok := shop["burger"]
	fmt.Println(burg, ok) // if ok is false it means variable isn't created

	fmt.Println("--------")
	fmt.Println("Before:")
	fmt.Println(shop)

	for k := range shop {
		shop[k]++
	}

	fmt.Println("--------")
	fmt.Println("After:")
	fmt.Println(shop)

	// 2
	weather := make(map[int]int)

	fmt.Println("---------")
	fmt.Println("Before:")
	fmt.Println(weather)
	fmt.Println("")

	weather[1] = 25
	weather[2] = 35
	weather[3] = 28

	fmt.Println("---------")
	fmt.Println("After:")
	fmt.Println(weather)
	fmt.Println("")
}
