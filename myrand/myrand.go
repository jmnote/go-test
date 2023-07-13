package myrand

import "math/rand"

func RandBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Dice() int {
	return rand.Intn(6) + 1
}

func Shuffle(nums []int) []int {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
