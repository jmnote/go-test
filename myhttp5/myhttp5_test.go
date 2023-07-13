package myhttp5

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
		start := time.Now()
		code, body := getAPIPing()
		elapsed := time.Since(start)
		assert.Equal(t, 200, code)
		assert.Equal(t, "pong", body)
		if i > 0 {
			assert.Greater(t, elapsed, time.Duration(300*time.Millisecond))
		}
	}
}
