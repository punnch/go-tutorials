package db

import (
	"company/company"
	"context"

	"github.com/jackc/pgx"
)

func SelectRows(ctx context.Context, pool *pgx.ConnPool) ([]company.Employee, error)
