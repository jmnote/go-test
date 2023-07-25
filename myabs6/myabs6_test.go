package myabs6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// three errors
func TestAbs_1(t *testing.T) {
	if Abs(-1) != -1 {
		t.Error("Abs(-1)")
	}
	if Abs(0) != -1 {
		t.Error("Abs(0)")
	}
	if Abs(1) != -1 {
		t.Error("Abs(1)")
	}
}

// Don't: not using t.Run
func TestAbs_1a(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int
	}{
		{"Abs(1)", 1, 9999}, // Abs(-1)
		{"Abs(2)", 2, 9999}, // Abs(0)
		{"Abs(3)", 3, 9999}, // Abs(1)
	}
	for _, tc := range testCases {
		got := Abs(tc.num)
		if got != tc.want {
			t.Errorf("%s = %d; want = %d", tc.name, got, tc.want)
		}
	}
}

func TestAbs_slice(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int
	}{
		{"Abs(1)", 1, 9999},
		{"Abs(2)", 2, 9999},
		{"Abs(3)", 3, 9999},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestAbs_map(t *testing.T) {
	testCases := map[string]struct {
		num  int
		want int
	}{
		"Abs(1)": {1, 9999},
		"Abs(2)": {2, 9999},
		"Abs(3)": {3, 9999},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestAbs_style1(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{1, 9999},
		{2, 9999},
		{3, 9999},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestAbs_style2(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{1, 9999}, // hello
		{2, 9999}, // hello#01
		{3, 9999}, // hello#02
	}
	for _, tc := range testCases {
		t.Run("hello", func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestAbs_style3(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int
	}{
		{"negative", -1, 9999},
		{"non-negative", 0, 9999},
		{"non-negative", 1, 9999},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestAbs_style4(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{-1, 9999},
		{0, 9999},
		{1, 9999},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Abs(%d)", i, tc.num), func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}
