package myerror3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNumber_ok(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"11", 11},
		{"-2", -2},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := ToNumber(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestToNumber_error(t *testing.T) {
	testCases := []struct {
		input     string
		want      int
		wantError string
	}{
		{"", 0, `strconv.Atoi: parsing "": invalid syntax`},
		{"a", 0, `strconv.Atoi: parsing "a": invalid syntax`},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := ToNumber(tc.input)
			assert.EqualError(t, err, tc.wantError)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestToNumber(t *testing.T) {
	testCases := []struct {
		input     string
		want      int
		wantError string
	}{
		{"11", 11, ""},
		{"-2", -2, ""},
		{"", 0, `strconv.Atoi: parsing "": invalid syntax`},
		{"a", 0, `strconv.Atoi: parsing "a": invalid syntax`},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := ToNumber(tc.input)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.wantError)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
