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

	fmt.Println("Guess the count from 1 to 100")
	fmt.Println("You have 10 attempts!")

	reader := bufio.NewReader(os.Stdin)

	for {
		if counter == 10 {
			fmt.Println("You didn't guess the count:(")
			fmt.Println("The guessed count -", random)
		}

		fmt.Println("------------------")
		fmt.Print("Input a number: ")

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
			fmt.Println("You guessed the count -", random, "!")
			fmt.Println("Attempts spent:", counter+1)
			return
		}

		if guess < random {
			counter += 1
			if counter != 10 {
				fmt.Println("The guessed count is greater")
				fmt.Println("Attempt count:", counter)
			}
			continue
		}

		if guess > random {
			counter += 1
			if counter != 10 {
				fmt.Println("The guessed count is less")
				fmt.Println("Attempt count:", counter)
			}
			continue
		}
	}
}
