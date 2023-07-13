package myunreachable

import (
	"encoding/json"
)

// myunreachable.go
type Data struct {
	Message string `json:"message"`
}

func JSONString(data Data) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err // unreachable
	}
	return string(bytes), nil
}
