package myhttpmock2

import (
	"fmt"
	"io"
	"net/http"
)

var endpoint = "http://example.com/api"

func Get(name string) (string, error) {
	// TODO: timeout
	resp, err := http.Get(endpoint + "/" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("err: %s", body)
	}
	return string(body), nil
}
