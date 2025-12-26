package main

import "fmt"

type Appliance interface {
	TurnOn()
}

// Coffe maker
type CoffeMaker string

func (c CoffeMaker) TurnOn() {
	fmt.Println("Coffe maker's turned on! (bhbhbhbh)")
}

// Fan
type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Fan's turned on! (owowooowow)")
}

// PC
type PC string

func (p PC) Play() {
	fmt.Println("'Playing minecraft'")
}

func (p PC) TurnOn() {
	fmt.Println("PC's turned on! (L;ASDKJFLK;ASDJF)")
}

func main() {
	var device Appliance

	device = CoffeMaker("PufiCoffe")
	device.TurnOn()

	device = Fan("Fascinated Fan")
	device.TurnOn()

	device = PC("i5-12400F, 2060S")
	device.TurnOn()
}
