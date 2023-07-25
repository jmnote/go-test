package mysquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	got := Square(-2)
	assert.Equal(t, 4, got)
}

func TestMultiply(t *testing.T) {
	got := multiply(-2, -2)
	assert.Equal(t, 4, got)
}
