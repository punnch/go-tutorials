package main

import (
	"fmt"
	"log"
)

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(i+1, "-", line)
	}
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Width %.2f can't be negative!", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("Height %.2f can't be negative!", height)
	}

	area := width * height

	return area / 10, nil
}

func tableTest() {
	// Таблица с Printf
	fmt.Printf("%12s | %s\n", "Name", "Amount")
	fmt.Println("-------------------------")
	fmt.Printf("%12s | %2d\n", "Paper", 10)
	fmt.Printf("%12s | %2d\n", "Apples", 30)
	fmt.Printf("%12s | %5.2f\n", "Bread", 2.37534)
}

func main() {
	// 1. tableTest
	tableTest()

	// 2. repeatLine
	repeatLine("Cumpot is the best", 5)

	// 3. paintNeeded
	var amount, total float64
	var err error

	amount, err = paintNeeded(5.3, 8.2) // 1 стена
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed for painting\n", amount)
	total += amount

	amount, err = paintNeeded(2.3, 8.2) // 2 стена
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed for painting\n", amount)
	total += amount

	fmt.Printf("Needed amount: %0.2f liters\n", total)
}
