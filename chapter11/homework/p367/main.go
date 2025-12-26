package main

import "fmt"

// Car type
type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

// Truck type
type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

func main() {
	var vehicle Vehicle

	// Car
	vehicle = Car("Toyoda Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")

	// Truck
	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}
