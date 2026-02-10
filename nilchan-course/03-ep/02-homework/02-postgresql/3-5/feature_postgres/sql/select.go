package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRow(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT id, title, author, review, year, read, created_at, read_at
	FROM books
	ORDER BY id ASC;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []BookModel
	for rows.Next() {
		var book BookModel

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Year,
			&book.Read,
			&book.CreatedAt,
			&book.ReadAt,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
