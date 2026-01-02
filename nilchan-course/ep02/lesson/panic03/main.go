package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Panic!", p)
		}
	}()

	slice := []int{1, 2, 5}
	fmt.Println(slice[3])
}
