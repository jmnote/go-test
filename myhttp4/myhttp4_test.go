package myhttp4

import (
	"bytes"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestRouter(t *testing.T) {
	assert.NotEmpty(t, router)
}

func TestAPILogin(t *testing.T) {
	testCase := []struct {
		username string
		password string
		wantCode int
		wantBody string
	}{
		{
			"", "",
			403, `{"error":"username or password is incorrect", "status":"error"}`,
		},
		{
			"hello", "",
			403, `{"error":"username or password is incorrect", "status":"error"}`,
		},
		{
			"hello", "foo",
			403, `{"error":"username or password is incorrect", "status":"error"}`,
		},
		{
			"", "world",
			403, `{"error":"username or password is incorrect", "status":"error"}`,
		},
		{
			"foo", "world",
			403, `{"error":"username or password is incorrect", "status":"error"}`,
		},
		{
			"hello", "world",
			200, `{"data":{"message":"logged in successfully", "token":"EXAMPLE"}, "status":"ok"}`,
		},
	}
	for _, tc := range testCase {
		t.Run("", func(t *testing.T) {
			w := httptest.NewRecorder()
			data := url.Values{}
			data.Add("username", tc.username)
			data.Add("password", tc.password)
			req := httptest.NewRequest("POST", "/api/login", bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.wantCode, w.Code)
			assert.JSONEq(t, tc.wantBody, w.Body.String())
		})
	}
}

func TestRun(t *testing.T) {
	go Run()
	time.Sleep(100 * time.Millisecond)
}
