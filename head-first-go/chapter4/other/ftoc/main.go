// Package ftoc converts fahrenheit  to celsius
package main

import (
	"fmt"
	"ftoc/keyboard"
	"log"
)

func main() {
	fmt.Print("Введите температуру в градусах Фарингейта: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// converts celsius to fahrenheits if no error
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Температура в градусах цельсия: %.2f", celsius)
}
