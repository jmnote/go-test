package myunreachable2b

import (
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
		{Data{Message: "hello"}, `{"message":"hello"}`},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			bytes, err := JSONString(tc.data, nil)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, bytes)
		})
	}
}

func TestJSONString_error(t *testing.T) {
	testCases := []struct {
		data      Data
		wantError string
	}{
		{Data{}, "marshal err: <nil> (fake)"},
		{Data{Message: "hello"}, "marshal err: <nil> (fake)"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := JSONString(tc.data, errors.New("fake"))
			assert.EqualError(t, err, tc.wantError)
			assert.Equal(t, "", got)
		})
	}
}
