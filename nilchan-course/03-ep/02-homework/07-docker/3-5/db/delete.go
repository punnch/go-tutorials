package db

import (
	"context"

	"github.com/jackc/pgx"
)

func DeleteRow(ctx context.Context, pool *pgx.ConnPool, id int) error
