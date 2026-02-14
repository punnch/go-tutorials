package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]User, error) {
	sql := `
	SELECT id, full_name, phone_number
	FROM users;
	`

	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.PhoneNumber,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
