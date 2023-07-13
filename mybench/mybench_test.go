package mybench

import (
	"math/big"
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt()
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(big.NewInt(0))
	}
}
