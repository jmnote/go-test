package mymock6

import (
	"testing"

	"github.com/jmnote/go-test/mymock6/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	m := mocks.NewICatClient(t)
	m.On("HTTPGet", "Alice").Return(200, "Hello Alice").Once()
	m.On("HTTPGet", "Bob").Return(200, "Bye Bob").Once()
	m.On("HTTPGet", "Carol").Return(403, "limit exceeded")

	catClient = m
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
