package myunreachable2

import (
	"encoding/json"
)

var jsonMarshal = json.Marshal

type Data struct {
	Message string `json:"message"`
}

func JSONString(data Data) (string, error) {
	bytes, err := jsonMarshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
