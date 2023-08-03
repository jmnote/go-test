package myunreachable2

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString_ok(t *testing.T) {
	testCases := []struct {
		data Data
		want string
	}{
		{Data{}, `{"message":""}`},
		{Data{Message: "foo"}, `{"message":"foo"}`},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			bytes, err := JSONString(tc.data)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, bytes)
		})
	}
}

func TestJSONString_error(t *testing.T) {
	jsonMarshal = func(v any) ([]byte, error) {
		return []byte{}, errors.New("fake")
	}
	testCases := []struct {
		data      Data
		wantError string
	}{
		{Data{}, "fake"},
		{Data{Message: "foo"}, "fake"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			str, err := JSONString(tc.data)
			assert.EqualError(t, err, tc.wantError)
			assert.Equal(t, "", str)
		})
	}
	jsonMarshal = json.Marshal
}
