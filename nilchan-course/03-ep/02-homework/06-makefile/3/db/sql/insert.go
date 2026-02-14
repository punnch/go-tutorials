package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, user User) error {
	sql := `
	INSERT INTO users (full_name, phone_number)
	VALUES($1, $2);
	`

	_, err := conn.Exec(ctx, sql, user.FullName, user.PhoneNumber)

	return err
}
