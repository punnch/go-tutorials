package library

import (
	"maps"
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

func (l *List) GetBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Book, len(l.books))

	maps.Copy(tmp, l.books)

	return tmp
}

func (l *List) GetAuthorBooks(author string) map[string]Book {
	authorBooks := make(map[string]Book)

	l.mtx.RLock()
	defer l.mtx.Unlock()

	for title, book := range l.books {
		if book.Author == author {
			authorBooks[title] = book
		}
	}

	return authorBooks
}

func (l *List) GetReadedBooks() map[string]Book {
	readedBooks := make(map[string]Book)

	l.mtx.RLock()
	defer l.mtx.RUnlock()

	for title, book := range l.books {
		if book.Readed {
			readedBooks[title] = book
		}
	}

	return readedBooks
}

func (l *List) GetUnreadBooks() map[string]Book {
	unreadBooks := make(map[string]Book)

	l.mtx.RLock()
	defer l.mtx.RUnlock()

	for title, book := range l.books {
		if !book.Readed {
			unreadBooks[title] = book
		}
	}

	return unreadBooks
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
