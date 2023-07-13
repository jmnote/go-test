package myabs

import (
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-2)
	if got != 2 {
		t.Errorf("Abs(-2) = %d; want 2", got)
	}
}
