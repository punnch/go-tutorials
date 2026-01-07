package main

import "fmt"

func main() {
	parking := map[string]float64{
		"A4": 400,
		"C4": 700,
		"B1": 1200,
		"F5": 250,
		"Y3": 650,
	}
	fmt.Println("------------------------------------")
	fmt.Println("Parking spaces with cost less than 500 rubles:")

	for i, v := range parking {
		if v < 500 {
			fmt.Println(i, v)
		}
	}
	fmt.Println("")
	fmt.Println("------------------------------------")
	fmt.Println("Parking spaces with cost greater than 900 rubles (10 percent discount)")

	for i, v := range parking {
		if v > 900 {
			v *= 0.9
			fmt.Println(i, v)
		}
	}
}
