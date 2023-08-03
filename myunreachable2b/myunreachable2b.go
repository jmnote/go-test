package myunreachable2b

import (
	"encoding/json"
	"fmt"
)

var fakeErr bool = false

type Data struct {
	Message string `json:"message"`
}

func JSONString(data Data) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil || fakeErr {
		return "", fmt.Errorf("marshal err: %v", err)
	}
	return string(bytes), nil
}
