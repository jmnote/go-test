package myskip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, Abs(-1), 1)
}

func TestSlowAbs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	assert.Equal(t, SlowAbs(-1), 1)
}
