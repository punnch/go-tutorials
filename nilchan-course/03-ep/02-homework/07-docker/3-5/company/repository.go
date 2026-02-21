package company

import (
	"context"
	"errors"
)

type WorkerRepository interface {
	InsertRow(ctx context.Context, employee Employee) (Employee, error)
	DeleteRow(ctx context.Context, id int) error
	SelectRows(ctx context.Context) ([]Employee, error)
}

var ErrNotFound = errors.New("employee not found")
