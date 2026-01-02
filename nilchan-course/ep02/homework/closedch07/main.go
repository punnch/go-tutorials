package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	interviewer := make(chan string)
	opinions := []string{"foolish", "nice:)", "idk", "good", "bad"}

	go func() {
		// range: 2-5
		peopleAmount := 2 + rand.IntN(4)
		fmt.Println("Interviewer met", peopleAmount, "people!")
		fmt.Println("Opinions:")

		// range: 300-700
		interval := 300 + rand.IntN(401)

		for i := 1; i <= peopleAmount; i++ {
			// range: 300-700ms
			time.Sleep(time.Duration(interval) * time.Millisecond)

			// get random opinion from a person
			randomOpinion := opinions[rand.IntN(5)]
			interviewer <- randomOpinion
		}

		close(interviewer)
	}()

	i := 1

	for opinion := range interviewer {
		fmt.Println(i, "-", opinion)
		i++
	}
}
