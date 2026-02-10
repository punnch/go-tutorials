package main

import (
	"context"
	"postgresql_study/feature_postgres/connection"
	"postgresql_study/feature_postgres/sql"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	now := time.Now()
	book := sql.BookModel{
		ID:        4,
		Title:     "Atomic habits",
		Author:    "James Clear",
		Review:    "awesome book",
		Year:      2018,
		Read:      true,
		CreatedAt: time.Now(),
		ReadAt:    &now,
	}

	// select
	books, err := sql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, v := range books {
		if v.ID == 4 {
			// update
			if err := sql.UpdateRow(ctx, conn, book); err != nil {
				panic(err)
			}
		}
	}
}
