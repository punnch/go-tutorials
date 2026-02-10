package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books  (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50) NOT NULL,
		author VARCHAR(50) NOT NULL,
		review VARCHAR(200),
		year INTEGER NOT NULL,
		read BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		read_at TIMESTAMP
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
