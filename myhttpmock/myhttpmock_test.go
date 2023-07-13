package myhttpmock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Alice")
	}))
	defer ts.Close()
	endpoint = ts.URL

	got, err := Get("Alice")
	assert.NoError(t, err)
	assert.Equal(t, "Hello Alice", got)
}
