package myhttp6

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestRouter(t *testing.T) {
	assert.NotEmpty(t, router)
}

func TestRun(t *testing.T) {
	go Run()
	time.Sleep(100 * time.Millisecond)
}

func getAPIPing() (int, string) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/ping", nil)
	router.ServeHTTP(w, req)
	return w.Code, w.Body.String()
}

func TestRateLimit(t *testing.T) {
	for i := 0; i < 5; i++ {
		code, body := getAPIPing()
		if i < 3 {
			assert.Contains(t, []int{200, 429}, code)
			assert.Contains(t, []string{"pong", "Too Many Requests"}, body)
		} else {
			assert.Equal(t, 429, code)
			assert.Equal(t, "Too Many Requests", body)
		}
	}
}
