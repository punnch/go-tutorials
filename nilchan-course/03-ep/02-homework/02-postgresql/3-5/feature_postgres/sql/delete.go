package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, IDs []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id=ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, IDs)

	return err
}
