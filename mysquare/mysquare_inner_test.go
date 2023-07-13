package mysquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	got := Square(-2)
	assert.Equal(t, got, 4)
}

func TestMultiply(t *testing.T) {
	got := multiply(-2, -2)
	assert.Equal(t, got, 4)
}
