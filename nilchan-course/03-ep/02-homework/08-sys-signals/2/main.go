package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", n, "done!")
			return
		case <-time.After(time.Second):
			fmt.Println("I'm worker", n, "working...")
		}
	}
}

func main() {
	countStr := os.Getenv("WORKERS_COUNT")
	if countStr == "" {
		fmt.Println("environment variable 'WORKERS_COUNT' has to be declared!")
		return
	}

	wg := &sync.WaitGroup{}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go worker(ctx, wg, i)
	}

	wg.Wait()

	fmt.Println("Program successfully ended!")
}
