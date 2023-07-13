// myabs2_test.go
package myabs2

import (
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-2)
	if got != 1 {
		t.Errorf("Abs(-2) = %d; want 2", got)
	}
}
