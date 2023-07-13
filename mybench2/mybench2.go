package mybench2

import (
	"time"
)

type Big struct {
}

func NewBig() *Big {
	time.Sleep(time.Second)
	return new(Big)
}

func (*Big) Len() int {
	return 42
}
