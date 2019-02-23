package golib

import (
	"math/big"
)

// Factorial returns n!
func Factorial(n int) string {
	ret := big.NewInt(1)
	for i := 1; i <= n; i++ {
		ret.Mul(ret, big.NewInt(int64(i)))
	}
	return ret.String()
}

// Fibonacci returns fibonacci
func Fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
