// myabs5_test.go
package myabs5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAbs_1(t *testing.T) {
	assert.Equal(t, 9999, Abs(1))
	assert.Equal(t, 9999, Abs(2))
	assert.Equal(t, 9999, Abs(3))
}

func TestAbs_2(t *testing.T) {
	require.Equal(t, 9999, Abs(1))
	require.Equal(t, 9999, Abs(2))
	require.Equal(t, 9999, Abs(3))
}
