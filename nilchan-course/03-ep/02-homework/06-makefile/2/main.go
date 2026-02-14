package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	countString := os.Getenv("OUTPUT_COUNT")
	if countString == "" {
		fmt.Println("environment variable 'OUTPUT_COUNT' has to be declared")
		return
	}

	outputCount, err := strconv.Atoi(countString)
	if err != nil {
		fmt.Println("failed to convert OUTPUT_COUNT into int")
		return
	}

	for i := 1; i <= outputCount; i++ {
		fmt.Println("Loop work:", i)
	}
}
