package mycover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	testCases := []struct {
		in   int
		want string
	}{
		{-1, "negative"},
		{5, "small"},
	}
	for _, tc := range testCases {
		got := Size(tc.in)
		assert.Equal(t, tc.want, got)
	}
}
