package myerror

import "strconv"

func ToNumber(str string) (int, error) {
	v, err := strconv.Atoi(str)
	return v, err
}