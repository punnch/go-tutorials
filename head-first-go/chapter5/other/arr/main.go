package main

import "fmt"

var arr = [5]int{1, 82, -5, 9, 0}

func basic() {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func ranged() {
	for _, value := range arr {
		fmt.Println(value)
	}
}

func main() {

}
