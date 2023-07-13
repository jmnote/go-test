package myhttpclient

import (
	"io"
	"net/http"
)

func Get(url string) (int, string, error) {
	resp, err := http.Get(url) // TODO: timeout
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	return resp.StatusCode, string(bodyBytes), nil
}
