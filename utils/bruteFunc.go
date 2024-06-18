package utils

import (
	"fmt"
	"time"
)

type fn func(arr []int) []int

func BruteFunc(f fn, size int) {
	arr := CreateSlice(size)
	for t := 0; t < 5; t++ {
		start := time.Now()
		_ = f(arr)
		end := time.Now()
		tDiff := end.Sub(start)
		fmt.Printf("Arr size: %d, Time: %s\n", size, tDiff)
		size *= 2
	}
}