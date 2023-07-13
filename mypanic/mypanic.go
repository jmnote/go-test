package mypanic

import (
	"errors"
	"fmt"
	"strconv"
)

func ToNumber(str string) (v int, err error) {
	defer myRecover(str, &err)
	v = toNumber(str)
	return v, nil
}

func toNumber(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(errors.New("not number"))
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
