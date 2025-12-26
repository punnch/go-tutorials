package main

import "fmt"

func main() {
	basicArr := [6]int{1, 81, 51, 39, 14, 5}
	slice := basicArr[3:6]

	// change arr values
	basicArr = [6]int{512, 58324, 50918, 5832, 412, 81}

	fmt.Println(basicArr)
	fmt.Println(slice)
}
