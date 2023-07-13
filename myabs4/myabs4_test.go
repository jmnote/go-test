// myabs4_test.go
package myabs4

import (
	"testing"
)

func TestAbs_1(t *testing.T) {
	if Abs(1) != 9999 {
		t.Error("Abs(1)")
	}
	if Abs(2) != 9999 {
		t.Error("Abs(2)")
	}
	if Abs(3) != 9999 {
		t.Error("Abs(3)")
	}
}

func TestAbs_2(t *testing.T) {
	if Abs(1) != 9999 {
		t.Fatal("Abs(1)")
	}
	if Abs(2) != 9999 {
		t.Fatal("Abs(2)")
	}
	if Abs(3) != 9999 {
		t.Fatal("Abs(3)")
	}
}
