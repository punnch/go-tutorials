package main

import (
	"context"
	"fmt"
	"time"
)

func parent(ctx context.Context, n int) {
	for {
		select{
		case <-ctx.Done():
			fmt.Println("parent function", n, "finished!")
			return 
		default:
			fmt.Println("parent", n)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func middle(ctx context.Context, n int) {
	for {
		select{
		case <-ctx.Done():
			fmt.Println("middle function", n, "finished!")
			return 	
		default:
			fmt.Println("middle", n)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func child(ctx context.Context, n int) {
	for {
		select{
		case <-ctx.Done():
			fmt.Println("child function", n, "finished!")
			return 	
		default:
			fmt.Println("child", n)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// context hierarchy
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)

	go parent(parentCtx, 1)
	go parent(parentCtx, 2)
	go parent(parentCtx, 3)

	go middle(middleCtx, 1)
	go middle(middleCtx, 2)
	go middle(middleCtx, 3)

	go child(childCtx, 1)
	go child(childCtx, 2)
	go child(childCtx, 3)

	time.Sleep(3 * time.Second)
	childCancel()

	time.Sleep(3 * time.Second)
	middleCancel()

	time.Sleep(3 * time.Second)
	parentCancel()

	time.Sleep(3 * time.Second)
	fmt.Println("I'm main function and I finished!")
}
