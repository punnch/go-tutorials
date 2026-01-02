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
	fmt.Println("Места с парковкой меньше 500 рублей:")

	for i, v := range parking {
		if v < 500 {
			fmt.Println(i, v)
		}
	}
	fmt.Println("")
	fmt.Println("------------------------------------")
	fmt.Println("Места с парковкой больше 900 рублей (10% скидка!):")

	for i, v := range parking {
		if v > 900 {
			v *= 0.9
			fmt.Println(i, v)
		}
	}
}
