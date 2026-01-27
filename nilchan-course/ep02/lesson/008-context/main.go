package main

import (
	"context"
	"fmt"
	"time"
)

// Parent context
func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo", n, "ended!")
			return
		default:
			fmt.Println("Foo", n)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

// Child context
func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo", n, "ended!")
			return
		default:
			fmt.Println("Boo", n)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go boo(childContext, 1)
	go boo(childContext, 2)
	go boo(childContext, 3)

	go foo(parentContext, 1)
	go foo(parentContext, 2)
	go foo(parentContext, 3)

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(1 * time.Second)
	parentCancel()

	time.Sleep(3 * time.Second)
	fmt.Println("main ended!")
}
