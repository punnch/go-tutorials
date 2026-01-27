package noteops

import (
	"study/notes"
)

func Validate(n notes.Note) bool {
	switch {
	case n.Title == "":
		return false
	case n.Message == "":
		return false
	default:
		return true
	}
}

func GetNoteById(id int) (notes.Note, bool) {
	for _, n := range notes.Messages {
		if n.ID == id {
			return n, true
		}
	}
	return notes.Note{}, false
}

func DeleteNoteById(id int) bool {
	for i, n := range notes.Messages {
		if n.ID == id {
			notes.Messages = append(notes.Messages[:i], notes.Messages[i+1:]...)
			return true
		}
	}
	return false
}
