package gridtraveler

// https://leetcode.com/problems/unique-paths/description/

var hashMap map[[2]int]int

func UniquePaths(m int, n int) int {
	if hashMap == nil {
		hashMap = make(map[[2]int]int)
	}
	if val, ok := hashMap[[2]int{m, n}]; ok {
		return val // Return stored result
	}
	if val, ok := hashMap[[2]int{n, m}]; { //
		return val 
	}
	if m == 1 || n == 1 {
		hashMap[[2]int{m,n}] = 1
		return 1
	}
	if m == 0 || n == 0 {
		hashMap[[2]int{m,n}] = 0
		return 0
	}

	result := UniquePaths(m-1, n) + UniquePaths(m, n-1)

	hashMap[[2]int{m, n}] = result

	return result
}

func UniquePathsCombinatory(m int, n int) int {
	ans := 1
	for i := 1; i <= m-1; i++ {
			ans = ans * (n - 1 + i) / i
	}
	return ans
}