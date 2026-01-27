package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	arr := [5]User{
		User{
			Name:    "Serega",
			Rating:  9.0,
			Premium: false,
		},
		User{
			Name:    "Тоха",
			Rating:  5.5,
			Premium: true,
		},
		User{
			Name:    "Nazar",
			Rating:  3.1,
			Premium: true,
		},
		User{
			Name:    "German",
			Rating:  0,
			Premium: true,
		},
		User{
			Name:    "George",
			Rating:  8.4,
			Premium: false,
		},
	}

	// До
	fmt.Println("--------")
	fmt.Println("Before:")
	for i, v := range arr {
		pp.Println(i, v)
	}

	// После
	fmt.Println("--------")
	fmt.Println("After:")

	// 1 вариант
	for i := 0; i < len(arr); i++ {
		if arr[i].Premium {
			arr[i].Rating += 1
		}
		pp.Println(i, arr[i])
	}

	// 2 вариант
	for i, v := range arr {
		if v.Premium {
			arr[i].Rating += 1
			v.Rating += 1
		}
		pp.Println(i, v)
	}

	// 3 вариант
	for index, user := range arr {
		if user.Premium {
			arr[index].Rating += 1
		}
		pp.Println(index, arr[index])
	}
}
