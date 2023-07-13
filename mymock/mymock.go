package mymock

import (
	"fmt"

	"github.com/jmnote/go-test/mymock/catclient"
)

func Get(name string) (string, error) {
	code, body := catclient.HTTPGet(name)
	if code != 200 {
		return "", fmt.Errorf("err: %s", body)
	}
	return body, nil
}
