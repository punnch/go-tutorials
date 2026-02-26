package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func job(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Job done!")
			return
		case <-time.After(time.Second):
			fmt.Println("Job duration:", i, "sec")
			i++
		}
	}
}

func main() {
	fmt.Println("PID:", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	job(ctx)

	fmt.Println("Program ended successfully!")
}
