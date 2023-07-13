package myunreachable3

import (
	"fmt"
	"log"
)

func Greet(ok bool) {
	if !ok {
		log.Fatal("bye")
	}
	fmt.Println("hello")
}
