package main

import "fmt"

type Gallons float64

// Supports Stringer interface
func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gal", g)
}

type Liters float64

// Supports Stringer interface
func (g Liters) String() string {
	return fmt.Sprintf("%.2f L", g)
}

type Mililiters float64

// Supports Stringer interface
func (g Mililiters) String() string {
	return fmt.Sprintf("%.2f mL", g)
}
func main() {
	fmt.Println(Gallons(15.0050158))
	fmt.Println(Liters(3.34567))
	fmt.Println(Mililiters(229.0083123))
}
