package library

import (
	"sync"
)

/*
1. Add book in library
2. Mark book as readed
3. Get a book
4. Get library + filter
  - author
  - Readed
  - Not readed
5. Delete a book
*/

type List struct {
	books map[string]Book
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		books: make(map[string]Book),
	}
}

func (l *List) AddBook(book Book) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.books[book.Title]
	if ok {
		return Book{}, ErrBookAlreadyExist
	}

	l.books[book.Title] = book

	return book, nil
}

func (l *List) ReadBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	book.Read()

	l.books[title] = book

	return book, nil
}

func (l *List) UnreadBook(title string) (Book, error) {
	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	book.Unread()

	l.books[title] = book

	return book, nil
}

func (l *List) GetBook(title string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	return book, nil
}

func (l *List) GetBooks(author string, read *bool) map[string]Book {
	tmp := make(map[string]Book)

	l.mtx.RLock()
	defer l.mtx.RUnlock()

	for title, book := range l.books {
		if author != "" && book.Author != author {
			continue
		}

		if read != nil && book.isRead != *read {
			continue
		}

		tmp[title] = book
	}

	return tmp
}

func (l *List) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}

	delete(l.books, title)

	return nil
}
