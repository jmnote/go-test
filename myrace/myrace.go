package myrace

import (
	"fmt"
	"time"
)

func Count() {
	i := 0
	go func() {
		for {
			i++ // write
			time.Sleep(time.Second)
			fmt.Println("in", i)
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("out", i) // read
}

func Count2() {
	i := 0
	go func() {
		j := i // new variable
		for {
			j++ // write
			time.Sleep(time.Second)
			fmt.Println("in", j)
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("out", i) // read
}
