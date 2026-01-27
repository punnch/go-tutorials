package main

import "fmt"

func main() {
	// 1 int array
	intArray := [3]int{16, 26, -4}

	// 2 float array
	floatArray := [2]float64{3.25, 7.75}

	// 3 boolean array
	boolArray := [2]bool{true, false}

	// 4 массив string типа
	strArray := [3]string{"a", "b", "c"}

	// Operations:

	// 1.Gettings of element
	el := intArray[0]
	fmt.Println("index 0:", el)

	// 2. Output
	fmt.Println(floatArray)

	// 3. Change element
	fmt.Println("----")
	fmt.Println("Before:")
	fmt.Println(boolArray[0])
	fmt.Println("")

	boolArray[0] = false

	fmt.Println("----")
	fmt.Println("After:")
	fmt.Println(boolArray[0])
	fmt.Println("")
	// 4. Arrays + Loops

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
