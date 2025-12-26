package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func main() {
	var numbers []float64
	arguments := os.Args[1:]

	for _, argument := range arguments {
		num, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	fmt.Printf("Average is: %.2f", average(numbers...))
}
