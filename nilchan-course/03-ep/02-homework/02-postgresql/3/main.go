package main

import (
	"context"
	"fmt"
	"postgresql_study/feature_postgres/connection"
)

func main() {
	ctx := context.Background()

	_, err := connection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("succeed!")
}
