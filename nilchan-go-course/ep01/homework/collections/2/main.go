package main

import "fmt"

func main() {
	// 1 слайс int типа, созданный на месте
	intSlice := []int{3, -2, 4}

	// 2 слайс str типа без выделения памяти
	strSlice := make([]string, 0)

	// 3 слайс str типа с выделением памяти
	strSlice2 := make([]string, 0, 3)

	// 3 слайс float типа с созданием элементов со значениями по-умолчанию
	floatSlice := make([]float64, 2)

	// Операции:
	// 1. Добавление нового элемента
	floatSlice = append(floatSlice, 0.999)
	strSlice2 = append(strSlice2, "lasdkjflsa")

	// 2. Получение отдельного элемента
	el2 := intSlice[2]
	fmt.Println("2 индекс:", el2)

	// 3. Вывод на экран
	fmt.Println(strSlice)

	// 4. Изменение отдельного элемента
	el0 := intSlice[0]
	fmt.Println("0 индекс до:", el0)

	el0 = 100
	fmt.Println("0 индекс после:", el0)

	// 5. Циклы
	for i, _ := range intSlice {
		fmt.Println(i, intSlice[i])
	}
}
