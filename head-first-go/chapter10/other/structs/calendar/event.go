package calendar

import (
	"errors"
	"unicode/utf8"
)

// Struct
type Event struct {
	title string
	Date
}

// Set title method
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 20 {
		return errors.New("invalid title length")
	}
	e.title = title
	return nil
}

// Get title method
func (e *Event) Title() string {
	return e.title
}
