package main

import "fmt"

func main() {
	// 1 str-int map
	school := map[string]int{
		"Dima":   5,
		"Anton":  5,
		"German": 2,
		"Cumpot": 0,
	}

	// 2 str-bool map created via make()
	hasWork := make(map[string]bool)

	hasWork["Dima"] = true
	hasWork["Anton"] = true
	hasWork["German"] = true
	hasWork["Cumpot"] = false

	// Operations:
	// 1. Delete element
	fmt.Println("Before:", school)

	delete(school, "German")

	fmt.Println("After:", school)

	// 2. Check default value
	c, ok := hasWork["Cumpot"]
	if !ok {
		fmt.Println("No such name in database!")
		return
	}

	fmt.Println("Name is in database...")

	if c {
		fmt.Println("Name has true value, You work!")
	} else {
		fmt.Println("Name has false value, You don't work:(")
	}
}
