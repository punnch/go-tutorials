package main

import (
	"fmt"
	"log"

	"github.com/punnch/average/datafile"
)

func main() {
	sum := 0.0

	arr, err := datafile.GetFloats("datafile/data.txt")
	if err != nil {
		log.Fatal()
	}

	for _, value := range arr {
		sum += value
	}

	average := sum / float64(len(arr))

	fmt.Printf("Среднее массива: %.2f", average)
}
