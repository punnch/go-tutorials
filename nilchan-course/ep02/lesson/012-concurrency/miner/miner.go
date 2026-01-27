package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Miner imitates miner work
func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("I'm a miner:", n, "My work day is ended!")
			return
		default:
			fmt.Println("I'm a miner:", n, "Start working...")
			time.Sleep(time.Second)
			fmt.Println("I'm a miner:", n, "Got coal:", power)

			transferPoint <- power // send-only chan
			fmt.Println("I'm a miner:", n, "Brought coal:", power)
		}
	}
}

// MinerPool initializes miners
func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	// go func() needed to return channel immediately
	// and read values in real time
	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
