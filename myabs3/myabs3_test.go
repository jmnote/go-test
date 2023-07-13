// myabs3_test.go
package myabs3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	got := Abs(-2)
	assert.Equal(t, got, 1)
}
