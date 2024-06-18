package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jay13jay/metaPractice/mergesort"
)

var size int = 10000000
func main() {
	// ret := combosum.ComboSum([]int{2, 3, 6, 7}, 7)
	// ret := combosum.ComboSum([]int{2, 3, 5}, 0)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr = append(arr, int(rand.Intn(size)))
	}
	for t := 0; t < 5; t++ {
		start := time.Now()
		_ = mergesort.MS(arr)
		end := time.Now()
		tDiff := end.Sub(start)
		fmt.Printf("Arr size: %d, Time: %s\n", size, tDiff)
		size *= 2
	}
}