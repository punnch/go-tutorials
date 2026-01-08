package postman

import (
	"context"
	"fmt"
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

	for {
		select {
		case <-ctx.Done():
			fmt.Println("I'm a postman:", n, "My work day is ended!")
			return
		default:
			fmt.Println("I'm a postman:", n, "Start working...")
			time.Sleep(time.Second)
			fmt.Println("I'm a postman:", n, "Took a mail:", mail)

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
		go Postman(ctx, wg, mailTransfterPoint, i, postmanMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransfterPoint)
	}()

	return mailTransfterPoint
}

// postmanMail returns mail text
func postmanMail(mainNumber int) string {
	pmail := map[int]string{
		1: "Brawl stars news",
		2: "New video of Punnch",
		3: "Golang update",
		4: "Machine Learning cource discount",
		5: "AI will replace programmers",
	}

	if _, ok := pmail[mainNumber]; !ok {
		return "Breaking News"
	}

	return pmail[mainNumber]
}
