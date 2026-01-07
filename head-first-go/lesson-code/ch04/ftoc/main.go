// Package ftoc converts fahrenheit  to celsius
package main

import (
	"fmt"
	"ftoc/keyboard"
	"log"
)

func main() {
	fmt.Print("Input fahrenheit temperature: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// converts celsius to fahrenheits if no error
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Calcius temperature: %.2f", celsius)
}
