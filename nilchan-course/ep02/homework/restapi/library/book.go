package library

import "time"

type Book struct {
	title  string
	author string
	pages  int
	isRead bool

	addedAt  time.Time
	readedAt *time.Time
}

func NewBook(title string, author string, pages int) Book {
	return Book{
		title:  title,
		author: author,
		pages:  pages,
		isRead: false,

		addedAt:  time.Now(),
		readedAt: nil,
	}
}

func (b *Book) Read() {
	b.isRead = true

	readedAt := time.Now()
	b.readedAt = &readedAt
}

func (b *Book) Unread() {
	b.isRead = false
	b.readedAt = nil
}
