package dto

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
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

func ErrorJSON(w http.ResponseWriter, err error, code int) {
	errorDTO := NewErrorDTO(err.Error())
	http.Error(w, errorDTO.ToString(), code)
}

func ErrorCompareJSON(w http.ResponseWriter, err, target error, code int) {
	errorDTO := NewErrorDTO(err.Error())

	if errors.Is(err, target) {
		http.Error(w, errorDTO.ToString(), code)
	} else {
		http.Error(w, errorDTO.ToString(), http.StatusInternalServerError)
	}
}
