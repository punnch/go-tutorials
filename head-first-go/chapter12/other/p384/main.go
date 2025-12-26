package main

import (
	"fmt"
	"log"
	"os"

	"github.com/punnch/datafile"
)

func main() {
	numbers, err := datafile.GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Sum: %.2f\n", sum)
}
