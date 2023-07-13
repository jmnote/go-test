// myexample.go
package myexample

import (
	"fmt"
	"math/rand"
)

func Hello() {
	fmt.Println("hello")
}

func Salutations() {
	fmt.Println("hello")
	fmt.Println("goodbye")
}

func Shuffle(nums []int) []int {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
