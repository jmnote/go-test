package myreverse

import "time"

func Reverse(str string) (result string) {
	time.Sleep(time.Duration(len(str)*100) * time.Millisecond) // 길수록 지연
	for _, v := range str {
		result = string(v) + result
	}
	return
}
