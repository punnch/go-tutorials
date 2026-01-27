package main

import (
	"fmt"
	"math"
)

func MaxCount(numbers ...float64) float64 {
	max := math.Inf(-1) // -Inf

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(MaxCount(1, 2, 3))
}
