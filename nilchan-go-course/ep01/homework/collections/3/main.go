package main

import "fmt"

func main() {
	// 1 мапа str-int
	school := map[string]int{
		"Дима":   5,
		"Антон":  5,
		"Герман": 2,
		"Кампот": 0,
	}

	// 2 мапа str-bool через make
	hasWork := make(map[string]bool)

	hasWork["Дима"] = true
	hasWork["Антон"] = true
	hasWork["Герман"] = true
	hasWork["Кампот"] = false

	// Операции:
	// 1. Удаление элементов
	fmt.Println("До:", school)

	delete(school, "Герман")

	fmt.Println("После:", school)

	// 2. Является ли полученное значение значеним по-умолчанию?
	c, ok := hasWork["Кампот"]
	if !ok {
		fmt.Println("Имени нет в базе данных!")
		return
	}

	fmt.Println("Имя есть в базе данных...")

	if c {
		fmt.Println("Имя имеет true значение, Вы работаете!")
	} else {
		fmt.Println("Имя имеет false значение, Вы не работаете(")
	}
}
