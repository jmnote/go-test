package myabs7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs_2depth(t *testing.T) {
	t.Run("foo", func(t *testing.T) {
		assert.Equal(t, Abs(-1), 1)
	})
	t.Run("bar", func(t *testing.T) {
		assert.Equal(t, Abs(-2), 2)
	})
}

func TestAbs_3depth(t *testing.T) {
	t.Run("foo", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, Abs(-11), 11)
		})
		t.Run("", func(t *testing.T) {
			assert.Equal(t, Abs(-12), 12)
		})
	})
	t.Run("bar", func(t *testing.T) {
		assert.Equal(t, Abs(-2), 2)
	})
}
