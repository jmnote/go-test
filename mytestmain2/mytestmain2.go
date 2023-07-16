package mytestmain2

import (
	"fmt"
	"io"
	"net/http"
)

var host = "http://example.com:9090"

func GetMetadata() (string, error) {
	resp, err := http.Get(host + "/api/v1/metadata")
	if err != nil {
		return "", fmt.Errorf("http.Get err: %e", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll err: %e", err)
	}
	return string(bodyBytes), nil
}
