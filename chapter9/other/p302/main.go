package main

import "fmt"

type dollars float64
type rubles float64

func main() {
	var dima dollars
	var german rubles

	// converts ru to dol
	dima = dollars(rubles(500) / 81.38)
	// converts dol to ru
	german = rubles(dollars(30) * 81.38)

	fmt.Printf("Dima: %.2f \nGerman: %.f\n", dima, german)
}
