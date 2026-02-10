package main

import (
	"context"
	"postgresql_study/feature_postgres/connection"
	"postgresql_study/feature_postgres/sql"
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

	if err := sql.ListPages(ctx, conn, 5); err != nil {
		panic(err)
	}
}
