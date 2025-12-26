package main

import (
	"fmt"
	"log"
)

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(i+1, "-", line)
	}
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Ширина %.2f не может быть отрицательной!", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("Высота %.2f не может быть отрицательной!", height)
	}

	area := width * height

	// допустим 1 литра краски хватает на 10м
	return area / 10, nil
}

func tableTest() {
	// Таблица с Printf
	fmt.Printf("%12s | %s\n", "Наименование", "Количество")
	fmt.Println("-------------------------")
	fmt.Printf("%12s | %2d\n", "Бумага", 10)
	fmt.Printf("%12s | %2d\n", "Скрепки", 30)
	fmt.Printf("%12s | %5.2f\n", "Хлеб", 2.37534)
}

func main() {
	// 1. tableTest
	tableTest()

	// 2. repeatLine
	repeatLine("Кампот лох", 5)

	// 3. paintNeeded
	var amount, total float64
	var err error

	amount, err = paintNeeded(5.3, 8.2) // 1 стена
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f литров нужно для покраски\n", amount)
	total += amount

	amount, err = paintNeeded(2.3, 8.2) // 2 стена
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f литров нужно для покраски\n", amount)
	total += amount

	fmt.Printf("Всего необходимо %0.2f литров\n", total)
}
