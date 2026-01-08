package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var coal int
	var mails []string

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

	coalTransferPoint := miner.MinerPool(minerCtx, 2)
	mailTransfterPoint := postman.PostmanPool(postmanCtx, 2)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal += v
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransfterPoint {
			mails = append(mails, v)
		}
	}()

	wg.Wait()

	fmt.Println("--------------")
	fmt.Println("Coal:", coal)
	fmt.Println("Mails:", len(mails))
}
