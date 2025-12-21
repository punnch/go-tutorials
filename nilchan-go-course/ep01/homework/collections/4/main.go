package main

import "fmt"

type Dog struct {
	// 1. Имя собаки
	Name string

	// 2. Рейтинг от судий собаки
	Rating float64

	// 3. Уважение к судьям
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
	fmt.Println("До оценки:", dogArr)
	fmt.Println("")

	for i := 0; i < len(dogArr); i++ {
		if dogArr[i].Rep {
			dogArr[i].Rating++
		}
	}

	fmt.Println("------------")
	fmt.Println("После:", dogArr)
	fmt.Println("")
}
