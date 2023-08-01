package myerror3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	testCases := []struct {
		input     string
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
			got, err := ToInt(tc.input)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.wantError)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
