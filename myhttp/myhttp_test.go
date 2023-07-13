package myhttp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPing_1(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPing_2(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestRun(t *testing.T) {
	go Run()
	time.Sleep(100 * time.Millisecond)
}

func TestRun_2(t *testing.T) {
	go Run()
	time.Sleep(100 * time.Millisecond)
	resp, err := http.Get("http://127.0.0.1:8080/ping")
	assert.NoError(t, err)
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(body), "pong")
}
