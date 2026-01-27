package main

import (
	"restapi/httpwork"
	"restapi/library"
)

func main() {
	bookList := library.NewList()
	httpHandlers := httpwork.NewHttpHandlers(bookList)
	httpServer := httpwork.NewHttpServer(httpHandlers)

	httpServer.Start()
}
