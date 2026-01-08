package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string
	var mtx sync.Mutex

	minerCtx, minerCancel := context.WithCancel(context.Background())
	postmanCtx, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
		fmt.Println("--> MINER WORKDAY IS OVER!!! <--")
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
		fmt.Println("--> POSTMAN WORKDAY IS OVER!!! <--")
	}()

	coalTransferPoint := miner.MinerPool(minerCtx, 6)
	mailTransfterPoint := postman.PostmanPool(postmanCtx, 6)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransfterPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("-----------------------")
	fmt.Println("Coal:", coal.Load())

	mtx.Lock()
	fmt.Println("Mails:", len(mails))
	mtx.Unlock()
}
