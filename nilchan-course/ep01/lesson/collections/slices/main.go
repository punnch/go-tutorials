package main

import "fmt"

func main() {
	strSlice := make([]string, 0, 5)

	fmt.Println("Before:")
	fmt.Println("len:", len(strSlice))
	fmt.Println("cap:", cap(strSlice))

	strSlice = append(strSlice, "first")
	strSlice = append(strSlice, "second")
	strSlice = append(strSlice, "third")

	fmt.Println("After:")
	fmt.Println("len:", len(strSlice))
	fmt.Println("cap:", cap(strSlice))

	fmt.Println(strSlice)
}
