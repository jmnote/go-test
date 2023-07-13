package myhttpmock2

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/Alice", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Alice")
	})
	mux.HandleFunc("/Bob", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Bob")
	})
	ts := httptest.NewUnstartedServer(mux)
	ts.Start()
	defer ts.Close()
	endpoint = ts.URL

	got, err := Get("Alice") // TODO: table-driven
	assert.NoError(t, err)
	assert.Equal(t, "Hello Alice", got)
	got, err = Get("Bob")
	assert.NoError(t, err)
	assert.Equal(t, "Hello Bob", got)
	got, err = Get("Carol")
	assert.EqualError(t, err, "err: 404 page not found\n")
	assert.Equal(t, "", got)
}
