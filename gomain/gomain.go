package main

import (
	"fmt"

	"github.com/juagargi/go-from-c/lib/golib"
)

func manyCalls(count int, argument int) {
	fmt.Printf("Calling ... ")
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", i+1)
		golib.Fibonacci(argument)
	}
	fmt.Printf(". Done.\n")
}

func main() {
	fmt.Printf("Hello world\n")
	var n = 32
	// fmt.Printf("%d!=%s\n", n, golib.Factorial(n))
	fmt.Printf("fib(%d)=%d\n", n, golib.Fibonacci(n))
	manyCalls(1000, n)
}
