package fib

import "fmt"

// Fibonacci returns the n-th fibonacci number.
//
// https://elearningindustry.com/fibonacci-sequence-what-is-and-how-applies-agile-development
func Fibonacci(n uint) (uint, error) {
	if n > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}

	var a, b uint = 0, 1
	for i := 0; uint(i) < n; i++ {
		a, b = b, a+b
	}

	return a, nil
}
