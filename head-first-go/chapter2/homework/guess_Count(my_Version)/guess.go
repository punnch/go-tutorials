package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	random := rand.Intn(100) + 1
	var counter int
	fmt.Println(random)

	fmt.Println("Я загадал число от 1 до 100")
	fmt.Println("У тебя есть 10 попыток, чтобы отгадать")

	reader := bufio.NewReader(os.Stdin)

	for {
		if counter == 10 {
			fmt.Println("К сожалению, ты не угадал число :(")
			fmt.Println("Загаданное число -", random)
		}

		fmt.Println("------------------")
		fmt.Print("Введите число: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == random {
			fmt.Println("Ты угадал, загаданное число -", random, "!")
			fmt.Println("Попыток потрачено:", counter+1)
			return
		}

		if guess < random {
			counter += 1
			if counter != 10 {
				fmt.Println("Загаданное число больше")
				fmt.Println("Число попыток:", counter)
			}
			continue
		}

		if guess > random {
			counter += 1
			if counter != 10 {
				fmt.Println("Загаданное число меньше")
				fmt.Println("Число попыток:", counter)
			}
			continue
		}
	}
}
