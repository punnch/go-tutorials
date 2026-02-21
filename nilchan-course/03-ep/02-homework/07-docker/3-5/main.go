package main

import (
	"context"
	"fmt"
	"os"

	"company/company"
	"company/db"
	"company/http_server"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	connString := os.Getenv("DB_URL")
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}

	repo := db.NewPostgresRepo(pool)
	enterprise := company.NewCompany(repo)
	handlers := http_server.NewHandlers(enterprise)
	server := http_server.NewServer(handlers)

	if err := server.Start(); err != nil {
		fmt.Println("error occurred during the server work:", err)
	}
}
