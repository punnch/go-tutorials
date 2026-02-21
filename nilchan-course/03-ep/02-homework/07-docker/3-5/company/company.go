package company

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Company struct {
	repo WorkerRepository
}

func NewCompany(repo WorkerRepository) *Company {
	return &Company{
		repo: repo,
	}
}

func (c *Company) CreateEmployee(ctx context.Context, fullName, position string) (Employee, error) {
	employee := NewEmployee(fullName, position)

	employeeDB, err := c.repo.InsertRow(ctx, employee)
	if err != nil {
		return Employee{}, err
	}

	return employeeDB, nil
}

func (c *Company) GetEmployees(ctx context.Context) ([]Employee, error) {
	employees, err := c.repo.SelectRows(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (c *Company) DeleteEmployee(ctx context.Context, id int) error {
	err := c.repo.DeleteRow(ctx, id)

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	return err
}
