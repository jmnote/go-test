package mysquare_test

import (
	"testing"

	"github.com/jmnote/go-test/mysquare"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	got := mysquare.Square(-2)
	assert.Equal(t, got, 4)
}

// func TestMultiply(t *testing.T) {
// 	got := mysquare.multiply(-2, -2)
// 	assert.Equal(t, got, 4)
// }
