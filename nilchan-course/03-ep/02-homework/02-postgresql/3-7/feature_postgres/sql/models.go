package sql

import "time"

type BookModel struct {
	ID     int
	Title  string
	Author string
	Review string
	Year   int
	Read   bool

	CreatedAt time.Time
	ReadAt    *time.Time
}
