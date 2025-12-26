package main

import "fmt"

type part struct {
	description string
	count       int
}

func doublePack(p *part) {
	p.count *= 2
}
func main() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
}
