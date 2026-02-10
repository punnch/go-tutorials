package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func ListPages(ctx context.Context, conn *pgx.Conn, n int) error {
	sqlQuery := `
	SELECT id, title, author, review, year, read, created_at, read_at
	FROM books
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	offset := 0
	pageNum := 1

	for {
		var books []BookModel

		rows, err := conn.Query(ctx, sqlQuery, n, offset)
		if err != nil {
			return err
		}
		defer rows.Close()

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
				return err
			}

			books = append(books, book)
		}

		if len(books) == 0 {
			return nil
		}

		pp.Printf("Page %d: %v\n", pageNum, books)

		if len(books) < n {
			break
		}

		offset += n
		pageNum++
	}

	return nil
}

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
