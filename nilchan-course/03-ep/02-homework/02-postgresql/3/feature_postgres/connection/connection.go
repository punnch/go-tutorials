package connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/nilchan-homework")
	if err != nil {
		return conn, err
	}

	err = conn.Ping(ctx)
	return conn, err
}
