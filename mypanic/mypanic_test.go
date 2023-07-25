package mypanic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToNumber(t *testing.T) {
	testCases := []struct {
		str       string
		want      int
		wantError string
	}{
		{"11", 11, ""},
		{"-2", -2, ""},
		{"", 0, "not number"},
		{"a", 0, "not number"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := ToNumber(tc.str)
			assert.Equal(t, tc.want, got)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, "not number")
			}
		})
	}
}

func Test_toNumber(t *testing.T) {
	testCases := []struct {
		str       string
		want      int
		wantPanic string
	}{
		{"11", 11, ""},
		{"-2", -2, ""},
		{"", 0, "not number"},
		{"a", 0, "not number"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if tc.wantPanic == "" {
				got := toNumber(tc.str)
				assert.Equal(t, tc.want, got)
			} else {
				assert.PanicsWithError(t, tc.wantPanic, func() {
					got := toNumber(tc.str)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
