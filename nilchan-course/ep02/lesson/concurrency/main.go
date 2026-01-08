package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"time"
)

func main() {
	var coal int
	var mails []string

	minerCtx, minerCancel := context.WithCancel(context.Background())
	postmanCtx, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerCtx, 2)
	mailTransfterPoint := postman.PostmanPool(postmanCtx, 2)

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()

	isMinerClosed := false
	isPostmanClosed := false

	for !isPostmanClosed || !isMinerClosed {
		select {
		case c, ok := <-coalTransferPoint:
			if !ok {
				isMinerClosed = true
				continue
			}
			coal += c
		case m, ok := <-mailTransfterPoint:
			if !ok {
				isPostmanClosed = true
				continue
			}
			mails = append(mails, m)
		}
	}
}
