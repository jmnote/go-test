package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jmnote/go-test/myabs"
)

func main() {
	numPtr := flag.String("num", "", "a string")
	flag.Parse()
	num, err := strconv.Atoi(*numPtr)
	if err != nil {
		fmt.Printf("error: %s is not number\n", *numPtr)
		os.Exit(1)
	}
	fmt.Printf("Abs(%d)=%d\n", num, myabs.Abs(num))
}
