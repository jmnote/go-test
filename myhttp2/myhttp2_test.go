package myhttp2

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := setupRouter()
	testCases := []struct {
		method   string
		path     string
		wantCode int
		wantBody string
	}{
		{
			"GET", "/v1/hello",
			404, "404 page not found",
		},
		{
			"POST", "/v1/login",
			200, "pong",
		},
		{
			"POST", "/v1/submit",
			202, "pong",
		},
		{
			"POST", "/v1/read",
			200, "pong",
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, tc.path, nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.wantCode, w.Code)
			assert.Equal(t, tc.wantBody, w.Body.String())
		})
	}
}
