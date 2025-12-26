package main

import (
	"fmt"
	"log"
)

type OverheatError float64

// Supports Error interface
func (o OverheatError) Error() string {
	return fmt.Sprintf("It's overheating by %.2f degrees temperature!", o)
}

func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe
	// if it's overheating, returns error
	if excess > 0 {
		return OverheatError(excess) // error type
	}
	return nil
}

func main() {
	var err error
	err = CheckTemperature(120, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It's not overheat temp!")
}
