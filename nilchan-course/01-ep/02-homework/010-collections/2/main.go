package main

import "fmt"

func main() {
	// 1 int slice
	intSlice := []int{3, -2, 4}

	// 2 string slice without memmory allocation
	strSlice := make([]string, 0)

	// 3 string slice with memory allocation
	strSlice2 := make([]string, 0, 3)

	// 4 string slice with default values
	floatSlice := make([]float64, 2)

	// Operations:
	// 1. Add new element
	floatSlice = append(floatSlice, 0.999)
	strSlice2 = append(strSlice2, "lasdkjflsa")

	// 2. Get element
	el2 := intSlice[2]
	fmt.Println("2 индекс:", el2)

	// 3. Output
	fmt.Println(strSlice)

	// 4. Change element
	el0 := intSlice[0]
	fmt.Println("0 индекс до:", el0)

	el0 = 100
	fmt.Println("0 индекс после:", el0)

	// 5. Loops
	for i, _ := range intSlice {
		fmt.Println(i, intSlice[i])
	}
}
