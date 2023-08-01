package mypanic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	testCases := []struct {
		str       string
		want      int
		wantError string
	}{
		{"11", 11, ""},
		{"-2", -2, ""},
		{"", 0, "not int"},
		{"a", 0, "not int"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := ToInt(tc.str)
			assert.Equal(t, tc.want, got)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, "not int")
			}
		})
	}
}

func Test_toInt(t *testing.T) {
	testCases := []struct {
		str       string
		want      int
		wantPanic string
	}{
		{"11", 11, ""},
		{"-2", -2, ""},
		{"", 0, "not int"},
		{"a", 0, "not int"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if tc.wantPanic == "" {
				got := toInt(tc.str)
				assert.Equal(t, tc.want, got)
			} else {
				assert.PanicsWithError(t, tc.wantPanic, func() {
					toInt(tc.str)
				})
			}
		})
	}
}
