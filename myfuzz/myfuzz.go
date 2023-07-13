package myfuzz

import (
	"errors"
	"strings"
)

func Foo(i int, s string) (string, error) {
	return "", nil
}

func DontSayBye(s string) error {
	if strings.Contains(s, "bye") {
		return errors.New("you said bye")
	}
	return nil
}

func DontSayGoodbye(s string) error {
	if strings.Contains(s, "goodbye") {
		return errors.New("you said goodbye")
	}
	return nil
}

func Abs(num int) int {
	if num == 999 {
		return -1
	}
	if num < 0 {
		return -num
	}
	return num
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Reverse2(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
