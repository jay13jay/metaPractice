package fib

import "fmt"

func CalcFib(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	// O(N^2) naiive implemtation
	fmt.Printf("calculating for %d and %d\n", n-1, n-2)
	return CalcFib(n-1) + CalcFib(n-2)
}

