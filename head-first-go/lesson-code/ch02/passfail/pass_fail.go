package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	// Сохраняет значение строки до определенного символа
	input, err := reader.ReadString('\n') // Сохраняет все до перехода на новую строку
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // Удаляет отступы в начале и в конче строки (/n в нашем случае)
	grade, err := strconv.ParseFloat(input, 64) // Преобразовывает значение строки во float64
	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60 {
		status = "passsed"
	} else {
		status = "failed"
	}

	fmt.Println("Grade of", grade, "is", status)
}
