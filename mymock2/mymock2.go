package mymock2

import (
	"fmt"

	"github.com/jmnote/go-test/mymock2/catclient"
)

var httpGet = catclient.HTTPGet

func Get(name string) (string, error) {
	code, body := httpGet(name)
	if code != 200 {
		return "", fmt.Errorf("err: %s", body)
	}
	return body, nil
}
