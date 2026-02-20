package main

import (
	"company/company"
	"company/http_server"
	"fmt"
)

func main() {
	company := company.NewCompany()
	handlers := http_server.NewHandlers(company)
	server := http_server.NewServer(handlers)

	if err := server.Start(); err != nil {
		fmt.Println("error occurred during the server work:", err)
	}
}
