package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, num := range numbers {
		if num >= min && num <= max {
			result = append(result, num)
		}
	}

	if len(result) == 0 {
		return result
	}
	return result
}

func main() {
	fmt.Println(inRange(0, 20, 30.5, 10.4, -1, 1, 0))
}
