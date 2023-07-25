package mymock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
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
			assert.Equal(t, tc.want, got)
		})
	}
}
