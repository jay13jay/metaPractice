package utils

import "math/rand"

// CreateSlice takes size as an argument and returns a slice with capacity of "size"
// filled with random integers between 0 and 99.
func CreateSlice(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
