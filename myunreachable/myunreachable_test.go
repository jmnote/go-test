package myunreachable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	testCases := []struct {
		data Data
		want string
	}{
		{Data{}, `{"message":""}`},
		{Data{Message: "hello"}, `{"message":"hello"}`},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			bytes, err := JSONString(tc.data)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, bytes)
		})
	}
}
