package main

import "fmt"

type Dog struct {
	// 1. Dog name
	Name string

	// 2. Dog rating
	Rating float64

	// 3. Respect to others
	Rep bool
}

func main() {
	dogArr := [3]Dog{
		Dog{
			Name:   "Bobik",
			Rating: 6.0,
			Rep:    true,
		},

		Dog{
			Name:   "Dick",
			Rating: 3.5,
			Rep:    false,
		},

		Dog{
			Name:   "Bob",
			Rating: 10.0,
			Rep:    true,
		},
	}

	fmt.Println("------------")
	fmt.Println("Before evaluation:", dogArr)
	fmt.Println("")

	for i := 0; i < len(dogArr); i++ {
		if dogArr[i].Rep {
			dogArr[i].Rating++
		}
	}

	fmt.Println("------------")
	fmt.Println("After:", dogArr)
	fmt.Println("")
}
