package httpwork

import (
	"encoding/json"
	"time"
)

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
