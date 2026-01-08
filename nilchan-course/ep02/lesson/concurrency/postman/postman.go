package postman

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Postman imitates postman's work
func Postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	n int,
	mail string,
) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("I'm a postman:", n, "My work day is ended!")
		return
	default:
		for {
			fmt.Println("I'm a postman:", n, "Start working...")
			time.Sleep(time.Second)
			fmt.Println("I'm a postman:", n, "Went to the post office with mail:", mail)

			transferPoint <- mail
			fmt.Println("I'm a postman:", n, "Brought a mail:", mail)
		}
	}
}

// PostmanPool initializes postmans
func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransfterPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransfterPoint, i, postmanMail())
	}

	go func() {
		wg.Wait()
		close(mailTransfterPoint)
	}()

	return mailTransfterPoint
}

// postmanMail returns mail text
func postmanMail() string {
	pmail := map[int]string{
		0: "Brawl stars news",
		1: "New video of Punnch",
		2: "Golang update",
		3: "Machine Learning cource discount",
		4: "AI will replace programmers",
		5: "Breaking News",
	}

	randomMail := rand.IntN(6)

	return pmail[randomMail]
}
