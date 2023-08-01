package myerror2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		testCases := []struct {
			input string
			want  int
		}{
			{"11", 11},
			{"-2", -2},
		}
		for _, tc := range testCases {
			t.Run("", func(t *testing.T) {
				got, err := ToInt(tc.input)
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
	t.Run("error", func(t *testing.T) {
		testCases := []struct {
			input     string
			want      int
			wantError string
		}{
			{"", 0, "not int"},
			{"a", 0, "not int"},
		}
		for _, tc := range testCases {
			t.Run("", func(t *testing.T) {
				got, err := ToInt(tc.input)
				assert.EqualError(t, err, tc.wantError)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}
