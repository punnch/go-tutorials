package main

import (
	"fmt"
	"hfgo/chapter14/other/p438/sentence"
)

func main() {
	three := []string{"banana", "apple", "peach"}
	two := []string{"pen", "pencil"}
	one := []string{"laptop"}
	none := []string{}

	fmt.Println(sentence.JoinWithCommas(three))
	fmt.Println(sentence.JoinWithCommas(two))
	fmt.Println(sentence.JoinWithCommas(one))
	fmt.Println(sentence.JoinWithCommas(none))
}
