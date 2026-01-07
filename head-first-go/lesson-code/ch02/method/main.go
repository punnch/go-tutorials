package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 1 method example
	now := time.Now()
	year := now.Year()
	pastYear := year - 1

	fmt.Println("Current year:", year)
	fmt.Println("Past year:", pastYear)

	// 2 method instance
	words := "Bu##y, what're you #oing?"
	replacer := strings.NewReplacer("#", "d")
	fixed := replacer.Replace(words)

	fmt.Println(fixed)

}
