// удваивает переданное число, выводит его до умножения и после
package main

import "fmt"

func double(count *float64) {
	// удваиваем через указатель
	*count *= 2
}

func main() {
	count := 5.0

	fmt.Println("до:", count)

	// принимает адресс переменной
	double(&count)
	fmt.Println("после:", count)
}
