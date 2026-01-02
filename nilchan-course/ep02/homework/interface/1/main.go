package main

import (
	"study/figures"
	"study/figures/methods"
)

// 1. Каждая фигура имеет свои функции для площади и периметра
// - возвращает float64

// 2. Взаимодействие через модуль
// - в модуле интерфейс
// - у модуля функция для вывода площади и периметра

func main() {
	rect := methods.Rect{Height: 10, Width: 4}
	triangle := methods.Triangle{A: 5, B: 6, C: 7, H: 10}
	circle := methods.Circle{Radius: 4}

	figures.PrintShape(rect)
	figures.PrintShape(triangle)
	figures.PrintShape(circle)
}
