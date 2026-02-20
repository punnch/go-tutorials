package db

import (
	"company/company"
	"context"

	"github.com/jackc/pgx"
)

func InsertRow(ctx context.Context, pool *pgx.ConnPool) (company.Employee, error)
