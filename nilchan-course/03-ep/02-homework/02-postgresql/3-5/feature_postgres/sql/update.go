package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
	UPDATE books
	SET title=$1, author=$2, review=$3, year=$4, read=$5, created_at=$6, read_at=$7
	WHERE id=$8;
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
		&book.ID,
	)

	return err
}
