package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("PID:", os.Getpid())

	syssignal := make(chan os.Signal, 1)
	signal.Notify(syssignal, syscall.SIGINT)

	fmt.Println("Use CTRL+C")

	for v := range syssignal {
		fmt.Println("System signal:", v)
	}
}
