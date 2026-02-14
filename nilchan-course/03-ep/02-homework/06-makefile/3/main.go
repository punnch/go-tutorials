package main

import (
	"context"
	"os"
	"log"
	"errors"

	"github.com/k0kubun/pp/v3"

	"makefile/db/connection"
	"makefile/db/sql"
	"makefile/scanner"
)

func main() {
	ctx := context.Background()
	conn, err := connection.CreateConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newUser := os.Getenv("NEW_USER")

	switch newUser {
	case "":
		err := errors.New("environment variable 'NEW_USER' has to be declared")
		log.Fatal(err)
	case "YES":
		user, err := scanner.NewUserInput()
		if err != nil {
			log.Fatal(err)
		}

		if err := sql.InsertRow(ctx, conn, user); err != nil {
			log.Fatal(err)
		}
	case "NO":
		users, err := sql.SelectRows(ctx, conn)
		if err != nil {
			log.Fatal(err)
			return
		}

		for _, user := range users {
			pp.Println(user)
		}
	default:
		err := errors.New("environment variable 'NEW_USER' accepts only 'YES', 'NO' values")
		log.Fatal(err)
	}
}
