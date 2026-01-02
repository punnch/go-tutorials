package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	go func() {
		for {
			ch1 <- Message{
				Author: "German",
				Text:   "What's up!",
			}
			time.Sleep(3 * time.Second)
		}

	}()

	go func() {
		for {
			ch2 <- Message{
				Author: "Nazar",
				Text:   "yoo",
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case friend1 := <-ch1:
			fmt.Println("I've got a message from", friend1.Author, "with a text:", friend1.Text)
		case friend2 := <-ch2:
			fmt.Println("I've got a message from", friend2.Author, "with a text:", friend2.Text)
		}
	}
}
