// удваивает переданное число, выводит его до умножения и после
package main

import "fmt"

func double(count *float64) {
	// double throw pointer
	*count *= 2
}

func main() {
	count := 5.0

	fmt.Println("до:", count)

	// accepts varialbe address
	double(&count)
	fmt.Println("после:", count)
}
