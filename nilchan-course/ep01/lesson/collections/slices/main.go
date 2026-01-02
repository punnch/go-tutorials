package main

import "fmt"

func main() {
	strSlice := make([]string, 0, 5)

	fmt.Println("До:")
	fmt.Println("len:", len(strSlice))
	fmt.Println("cap:", cap(strSlice))

	strSlice = append(strSlice, "first")
	strSlice = append(strSlice, "second")
	strSlice = append(strSlice, "third")

	fmt.Println("После:")
	fmt.Println("len:", len(strSlice))
	fmt.Println("cap:", cap(strSlice))

	fmt.Println(strSlice)
}
