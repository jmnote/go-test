package pkg1

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	time.Sleep(time.Second)
	t.Fail()
}

func Test2(t *testing.T) {
	time.Sleep(time.Second)
	t.Fail()
}

func Test3(t *testing.T) {
	time.Sleep(time.Second)
	t.Fail()
}
