package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
	INSERT INTO books (title, author, review, year, read, created_at, read_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		&book.Title,
		&book.Author,
		&book.Review,
		&book.Year,
		&book.Read,
		&book.CreatedAt,
		&book.ReadAt,
	)

	return err
}
