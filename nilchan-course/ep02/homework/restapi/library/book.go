package library

import "time"

type Book struct {
	Title  string
	Author string
	Pages  int
	IsRead bool

	AddedAt  time.Time
	ReadedAt *time.Time
}

func NewBook(title string, author string, pages int) Book {
	return Book{
		Title:  title,
		Author: author,
		Pages:  pages,
		IsRead: false,

		AddedAt:  time.Now(),
		ReadedAt: nil,
	}
}

func (b *Book) Read() {
	b.IsRead = true

	readedAt := time.Now()
	b.ReadedAt = &readedAt
}

func (b *Book) Unread() {
	b.IsRead = false
	b.ReadedAt = nil
}
