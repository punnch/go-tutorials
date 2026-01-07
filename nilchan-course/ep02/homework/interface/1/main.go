package main

import (
	"study/figures"
	"study/figures/methods"
)

// 1. Each figure has functions Area() and Perim()
// - return float64

// 2. Module interaction
// - interface is in module
// - module has function to Print area and perimiter

func main() {
	rect := methods.Rect{Height: 10, Width: 4}
	triangle := methods.Triangle{A: 5, B: 6, C: 7, H: 10}
	circle := methods.Circle{Radius: 4}

	figures.PrintShape(rect)
	figures.PrintShape(triangle)
	figures.PrintShape(circle)
}
