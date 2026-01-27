package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	// Saves line value before certain symbol
	input, err := reader.ReadString('\n') // Saves all before new line
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // Deletes indentations
	grade, err := strconv.ParseFloat(input, 64) // Converts string value to float64
	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "passsed"
	} else {
		status = "failed"
	}

	fmt.Println("Grade of", grade, "is", status)
}
