package myerror

import (
	"errors"
	"strconv"
)

func ToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("not int")
	}
	return num, nil
}
