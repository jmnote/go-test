// myexample.go
package myexample

import (
	"fmt"
	"math/rand"
)

func Hello() {
	fmt.Println("hello")
}

func HelloBye() {
	fmt.Println("hello")
	fmt.Println("bye")
}

func Shuffle(nums []int) []int {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
