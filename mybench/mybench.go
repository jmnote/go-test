package mybench

import (
	"math/big"
	"math/rand"
)

func RandInt() int {
	return rand.Int()
}

func Factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, Factorial(n.Sub(x, n)))
}
