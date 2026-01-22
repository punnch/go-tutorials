package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// DTO - data transfer object

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type CompleteTaskDTO struct {
	Completed bool
}

func NewErrorDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// CompareSendErr compares two errors to check which encountered and sends error to client
func (e ErrorDTO) CompareSendErr(w http.ResponseWriter, err error, target error, code int) {
	if errors.Is(err, target) {
		http.Error(w, e.ToString(), code)
	} else {
		http.Error(w, e.ToString(), http.StatusInternalServerError)
	}
}
