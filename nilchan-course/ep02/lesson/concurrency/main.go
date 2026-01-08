package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
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
		fmt.Println("--> MINER WORKING DAY ENDED!!! <--")
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
		fmt.Println("--> POSTMAN WORKING DAY ENDED!!! <--")
	}()

	coalTransferPoint := miner.MinerPool(minerCtx, 2)
	mailTransfterPoint := postman.PostmanPool(postmanCtx, 2)

	isCoalClosed := false
	isMailClosed := false

	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}
			coal += c
		case m, ok := <-mailTransfterPoint:
			if !ok {
				isMailClosed = true
				continue
			}
			mails = append(mails, m)
		}
	}

	fmt.Println("--------------")
	fmt.Println("Coal:", coal)
	fmt.Println("Mails:", len(mails))
}
