package mergesort

import (
	"fmt"
)

// Merge sort 2 arrays
// takes an integer array as input
// returns the sorted array
func MS(arr []int) ([]int, error){
	var err error
	var ret []int
	var left, right []int
	if len(arr) < 2 {
		err = fmt.Errorf("error: cannot sort 1 element array")
		return arr, err
	}

	mid := len(arr) / 2
	left, _ = MS(arr[:mid])
	right, _ = MS(arr[mid:])

	ret = merge(left, right)
	return ret, err
}

func merge(left, right []int) []int {
	// l, r are indexes for left and right array
	// sort functions by sorting value 0 of each array,
	// then iterating the array that had the smaller value
	l, r := 0, 0 
	totLen := len(left) + len(right)
	result := make([]int, totLen)

	for i := 0; i < totLen; i++ {
		if l >= len(left) {
			result[i] = right[r]
			fmt.Printf("l > len(left) - setting result[%d] to right[%d]: %d\n", i, r, right[r])
			r++
		} else if r >= len(right) {
			result[i] = left[l]
			fmt.Printf("r > len(right) - setting result[%d] to left[%d]: %d\n", i, l, left[l])
			l++
		} else if left[l] < right[r] {
			fmt.Printf("left[l] < right[r] - setting result[%d] to left[%d]: %d\n", i, l, left[l])
			result[i] = left[l]
			l++
		} else {
			fmt.Printf("else - setting result[%d] to right[%d]: %d\n", i, r, right[r])
			result[i] = right[r]
			r++
		}

	}

	return result
}