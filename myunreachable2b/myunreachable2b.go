package myunreachable2b

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Message string `json:"message"`
}

var fakeErr error = nil

func JSONString(data Data) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil || fakeErr != nil {
		return "", fmt.Errorf("marshal err: %v (%v)", err, fakeErr)
	}
	return string(bytes), nil
}
