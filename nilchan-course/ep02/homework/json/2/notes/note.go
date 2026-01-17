package notes

import "time"

type Note struct {
	Title   string    `json:"header"`
	ID      int       `json:"id"` // user don't write that field
	Message string    `json:"message"`
	Urgent  bool      `json:"urgent"`
	Created time.Time `json:"created"` // user don't write that field
}

var (
	Messages []Note
	Id       int
)
