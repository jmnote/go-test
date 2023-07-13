package myskip

import "time"

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func SlowAbs(num int) int {
	time.Sleep(9 * time.Second)
	return Abs(num)
}
