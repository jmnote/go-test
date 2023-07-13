package myabs8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs_1(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{-1, 1},
		{-2, 2},
		{-3, 3},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Abs(%d)", i, tc.num), func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestAbs_2(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{0, 0},
		{1, 1},
		{100, 99},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d Abs(%d)", i, tc.num), func(t *testing.T) {
			got := Abs(tc.num)
			assert.Equal(t, got, tc.want)
		})
	}
}
