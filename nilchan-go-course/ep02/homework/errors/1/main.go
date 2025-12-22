package main

import (
	"fmt"
	"study/calculator"
)

func main() {
	myNums := calculator.Numbers{A: 1, B: 2}
	myNums2 := calculator.Numbers{A: 1, B: 0}

	// Sum
	num, err := calculator.Sum(myNums)
	if err != nil {
		fmt.Println("Sum error:", err)
	} else {
		fmt.Printf("Sum was successful!\nArguments: %d, %d\nResult: %d\n", myNums.A, myNums.B, num)
	}

	fmt.Println()

	// Multiplication
	num, err = calculator.Multiplicate(myNums)
	if err != nil {
		fmt.Println("Multiplication error:", err)
	} else {
		fmt.Printf("Multiplication was successful!\nArguments: %d, %d\nResult: %d\n", myNums.A, myNums.B, num)
	}

	fmt.Println()

	// Division by 0
	num, err = calculator.Divide(myNums2)
	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Printf("Division was successful!\nArguments: %d, %d\nResult: %d\n", myNums.A, myNums.B, num)
	}
}
