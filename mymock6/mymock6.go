package mymock6

import (
	"fmt"

	"github.com/jmnote/go-test/mymock6/catclient"
)

type ICatClient interface {
	HTTPGet(name string) (int, string)
}

var catClient = catclient.NewCatClient()

func Get(name string) (string, error) {
	code, body := catClient.HTTPGet(name)
	if code != 200 {
		return "", fmt.Errorf("err: %s", body)
	}
	return body, nil
}
