package mymock2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockHTTPGet(name string) (int, string) {
	switch name {
	case "Alice":
		return 200, "Hello Alice"
	case "Bob":
		return 200, "Bye Bob"
	case "Carol":
		return 403, "limit exceeded"
	}
	return 0, ""
}

func TestGet(t *testing.T) {
	httpGet = mockHTTPGet
	testCases := []struct {
		name      string
		want      string
		wantError string
	}{
		{"Alice", "Hello Alice", ""},
		{"Bob", "Bye Bob", ""},
		{"Carol", "", "err: limit exceeded"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := Get(tc.name)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.wantError)
			}
			assert.Equal(t, got, tc.want)
		})
	}
}
