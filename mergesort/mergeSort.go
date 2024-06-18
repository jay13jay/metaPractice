package mergesort

// Merge sort 2 arrays
// takes an integer array as input
// returns the sorted array
func MS(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := MS(arr[:mid])
	right := MS(arr[mid:])

	ret := merge(left, right)
	return ret
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
			r++
		} else if r >= len(right) {
			result[i] = left[l]
			l++
		} else if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
	}

	return result
}
