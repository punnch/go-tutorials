package dto

import "encoding/json"

type EployeeDTO struct {
	FullName string `json:"fullname"`
	Position string `json:"position"`
}

func ToJSON(v any) []byte {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}

	return b
}
