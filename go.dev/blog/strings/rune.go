package main

import "fmt"

func main() {
	// string is a slice of bytes
	const hey = "привет"
	for i, ch := range hey {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	} 

	fmt.Println(len(hey))
	fmt.Println(len("привет"))
}
