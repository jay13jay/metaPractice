package missingnum

// S = sum of the consecutive integers
// n = number of integers
// a = first term
// l = last term
// S = n(a + l)/2

func FindMissing(nums []int) int {
	a := 0
	l := len(nums)
	n := l + 1
	S := n * (a + l) / 2

	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	return sum - S
}