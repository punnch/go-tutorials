package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("PID:", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	fmt.Println("Until context is cancelled:")
	<-ctx.Done()
	fmt.Println("After context is cancelled.")
}
