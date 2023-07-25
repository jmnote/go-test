package myrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandBetween(t *testing.T) {
	for i := 0; i < 10000; i++ {
		got := RandBetween(1, 6)
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6}, got)
	}
}

func TestRandBetween_2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		got := RandBetween(1, 6)
		assert.GreaterOrEqual(t, got, 1)
		assert.LessOrEqual(t, got, 6)
	}
}

func TestRandBetween_3(t *testing.T) {
	min := 9999
	max := -9999
	for i := 0; i < 100000; i++ {
		got := RandBetween(1, 6)
		if got > max {
			max = got
		}
		if got < min {
			min = got
		}
	}
	assert.Equal(t, 1, min)
	assert.Equal(t, 6, max)
}

func TestDice(t *testing.T) {
	for i := 0; i < 10000; i++ {
		got := Dice()
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6}, got)
	}
}

func TestDice_2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		got := Dice()
		assert.GreaterOrEqual(t, got, 1)
		assert.LessOrEqual(t, got, 6)
	}
}

func TestDice_3(t *testing.T) {
	min := 9999
	max := -9999
	for i := 0; i < 100000; i++ {
		got := Dice()
		if got > max {
			max = got
		}
		if got < min {
			min = got
		}
	}
	assert.Equal(t, 1, min)
	assert.Equal(t, 6, max)
}

func TestShuffle(t *testing.T) {
	firstMap := map[int]bool{}
	for i := 0; i < 1000; i++ {
		got := Shuffle([]int{1, 2, 3, 4})
		assert.ElementsMatch(t, got, []int{1, 2, 3, 4})
		firstMap[got[0]] = true
	}
	assert.Equal(t, map[int]bool(map[int]bool{1: true, 2: true, 3: true, 4: true}), firstMap)
}
