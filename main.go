package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/mergesort"
)

func main() {
	// ret := combosum.ComboSum([]int{2, 3, 6, 7}, 7)
	// ret := combosum.ComboSum([]int{2, 3, 5}, 0)
	arr := []int{5,70,11,11,26,28,83,34,35,14,19,22}
	res := mergesort.MS(arr)

	fmt.Printf("\nResults:\n%v\n", res)
}