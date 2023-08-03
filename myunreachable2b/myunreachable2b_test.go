package myunreachable2b

import (
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
	fakeErr = true
	testCases := []struct {
		data      Data
		wantError string
	}{
		{Data{}, "marshal err: <nil>"},
		{Data{Message: "hello"}, "marshal err: <nil>"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := JSONString(tc.data)
			assert.EqualError(t, err, tc.wantError)
			assert.Equal(t, "", got)
		})
	}
	fakeErr = false
}
