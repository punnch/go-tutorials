package httpwork

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type BookReadDTO struct {
	isRead bool
}

type BookDTO struct {
	title  string
	author string
	pages  int
}

func (b BookDTO) VerifyToCreate() error {
	if b.title == "" {
		return errors.New("title is empty")
	}

	if b.author == "" {
		return errors.New("author is empty")
	}

	if b.pages == 0 {
		return errors.New("pages are empty")
	}

	return nil
}

type ErrorDTO struct {
	message string
	time    time.Time
}

func NewErrorDTO(msg string) ErrorDTO {
	return ErrorDTO{
		message: msg,
		time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ErrJSON(w http.ResponseWriter, err error, code int) {
	errorDTO := NewErrorDTO(err.Error())
	http.Error(w, errorDTO.ToString(), code)
}

func ErrCompareJSON(w http.ResponseWriter, err error, target error, code int) {
	errorDTO := NewErrorDTO(err.Error())

	if errors.Is(err, target) {
		http.Error(w, errorDTO.ToString(), code)
	} else {
		http.Error(w, errorDTO.ToString(), http.StatusInternalServerError)
	}
}

func ToJSON(v any) []byte {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	return b
}
