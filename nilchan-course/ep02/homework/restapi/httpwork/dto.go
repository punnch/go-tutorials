package httpwork

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type BookReadDTO struct {
	IsRead bool
}

type BookDTO struct {
	Title  string
	Author string
	Pages  int
}

func (b BookDTO) VerifyToCreate() error {
	if b.Title == "" {
		return errors.New("title is empty")
	}

	if b.Author == "" {
		return errors.New("author is empty")
	}

	if b.Pages == 0 {
		return errors.New("pages are empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(msg string) ErrorDTO {
	return ErrorDTO{
		Message: msg,
		Time:    time.Now(),
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
