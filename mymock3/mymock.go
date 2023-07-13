package mymock

import (
	"fmt"

	"github.com/jmnote/go-test/mymock3/catclient"
)

var catClient = catclient.NewCatClient()

func Get(name string) (string, error) {
	code, body := catClient.HTTPGet(name)
	if code != 200 {
		return "", fmt.Errorf("err: %s", body)
	}
	return body, nil
}
