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

	fmt.Println("Guess the count from 1 to 100")

	for i := 0; i < 10; i++ {
		fmt.Println("Attempt amount:", 10-i)
		fmt.Print("Input a number: ")

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
			fmt.Println("The guessed count is greater!")
		} else if guess > random {
			fmt.Println("The guessed count is less!")
		} else {
			success = true
			fmt.Println("You guessed!")
			break
		}
	}
	if !success {
		fmt.Println("No attempts :( The guessed count -", random)
	}
}
