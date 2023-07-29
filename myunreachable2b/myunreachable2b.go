package myunreachable2b

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Message string `json:"message"`
}

func JSONString(data Data, fakeErr error) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil || fakeErr != nil {
		return "", fmt.Errorf("marshal err: %v (%s)", err, fakeErr)
	}
	return string(bytes), nil
}
