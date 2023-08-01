package mypanic

import (
	"errors"
	"fmt"
	"strconv"
)

func ToInt(str string) (v int, err error) {
	defer myRecover(str, &err)
	v = toInt(str)
	return v, nil
}

func toInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(errors.New("not int"))
	}
	return v
}

func myRecover(str string, errp *error) {
	e := recover()
	if e == nil {
		return
	}
	*errp = fmt.Errorf("%v", e)
}
