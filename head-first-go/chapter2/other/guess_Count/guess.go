// guess - игра, в которой игроку нужно угадать загаданное рандомайзером число
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
	success := false

	// Тест
	fmt.Println(random)

	fmt.Println("Угадай число от 1 до 100")

	for i := 0; i < 10; i++ {
		fmt.Println("Количество попыток:", 10-i)
		fmt.Print("Введите число: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < random {
			fmt.Println("Загаданное число больше!")
		} else if guess > random {
			fmt.Println("Загаданное число меньше!")
		} else {
			success = true
			fmt.Println("Ты угадал!")
			break
		}
	}
	if !success {
		fmt.Println("Попытки кончились :( Загаданное число -", random)
	}
}
