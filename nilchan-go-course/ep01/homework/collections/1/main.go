package main

import "fmt"

func main() {
	// 1 массив int типа
	intArray := [3]int{16, 26, -4}

	// 2 массив float типа
	floatArray := [2]float64{3.25, 7.75}

	// 3 массив boolean типа
	boolArray := [2]bool{true, false}

	// 4 массив string типа
	strArray := [3]string{"a", "b", "c"}

	// Операции:

	// 1.Получение элемента
	el := intArray[0]
	fmt.Println("индекс 0:", el)

	// 2. Вывод на экран
	fmt.Println(floatArray)

	// 3. Изменение отдельного элемента
	fmt.Println("----")
	fmt.Println("До:")
	fmt.Println(boolArray[0])
	fmt.Println("")

	boolArray[0] = false

	fmt.Println("----")
	fmt.Println("После:")
	fmt.Println(boolArray[0])
	fmt.Println("")
	// 4. Массивы + циклы

	// 1)
	for i := 0; i < 3; i++ {
		fmt.Println(i, "-", strArray[i])
	}

	// 2)
	for i, _ := range intArray {
		if intArray[i]%2 == 0 {
			intArray[i] *= 2
		}
	}
	fmt.Println(intArray)
}
