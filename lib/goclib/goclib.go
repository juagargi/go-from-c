package main

import "C"
import "github.com/juagargi/go-from-c/lib/golib"

// Factorial calls factorial
func Factorial(n int) string {
	return golib.Factorial(n)
}

//export Fibonacci
func Fibonacci(n int) int {
	return golib.Fibonacci(n)
}

func main() {}
