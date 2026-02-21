package db

import (
	"company/company"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		pool: pool,
	}
}

func (p *PostgresRepo) InsertRow(ctx context.Context, employee company.Employee) (company.Employee, error) {
	sql := `
	INSERT INTO workers (full_name, position) 
	VALUES ($1, $2)
	RETURNING id, full_name, position;
	`

	var employeeDB company.Employee
	if err := p.pool.QueryRow(
		ctx,
		sql,
		employee.FullName,
		employee.Position,
	).Scan(
		&employeeDB.ID,
		&employeeDB.FullName,
		&employeeDB.Position,
	); err != nil {
		return company.Employee{}, err
	}

	return employeeDB, nil
}

func (p *PostgresRepo) DeleteRow(ctx context.Context, id int) error {
	sql := `
	DELETE FROM workers
	WHERE id=$1
	RETURNING id;
	`

	var checkID int
	err := p.pool.QueryRow(ctx, sql, id).Scan(&checkID)

	return err
}

func (p *PostgresRepo) SelectRows(ctx context.Context) ([]company.Employee, error) {
	sql := `
	SELECT id, full_name, position
	FROM workers;
	`

	var employees []company.Employee
	rows, err := p.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee company.Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Position,
		); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return employees, nil
}
