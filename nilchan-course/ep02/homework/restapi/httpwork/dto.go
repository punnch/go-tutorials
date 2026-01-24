package httpwork

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type TaskDTO struct {
	title  string
	author string
	pages  int
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
