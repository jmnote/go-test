package myreverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse_1(t *testing.T) {
	got := Reverse("hello")
	assert.Equal(t, "olleh", got)
}

func TestReverse_2(t *testing.T) {
	got := Reverse("foo")
	assert.Equal(t, "oof", got)
}

func TestReverseParallel_1(t *testing.T) {
	t.Parallel()
	got := Reverse("hello")
	assert.Equal(t, "olleh", got)
}

func TestReverseParallel_2(t *testing.T) {
	t.Parallel()
	got := Reverse("foo")
	assert.Equal(t, "oof", got)
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"foo", "oof"},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.input, func(t *testing.T) {
			got := Reverse(tc.input)
			t.Log(got)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestReverseParallel(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"foo", "oof"},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()
			got := Reverse(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
