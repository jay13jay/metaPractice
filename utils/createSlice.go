package utils

import "math/rand"

// Takes size as an argument
// Returns a slice with capacity of "size".
// Creates slice with rand.Intn(size)
func CreateSlice(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr = append(arr, int(rand.Intn(size)))
	}
	return arr
}