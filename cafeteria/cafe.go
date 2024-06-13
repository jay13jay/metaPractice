package cafeteria

import (
	"sort"
)

// S: taken seats
// N: Total seats
// M: # of current diners
// K: # of seats between each pair of adjacent seats

func GetMaxAdditionalDinersCount(N int64, K int64, M int64, S []int64) int64 {
  // Write your code here

	// S = the currently taken seats
	// N = total number of seats
	// K = required gap between seats
	// M = number of current diners

	var ret int64 = 0
	sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })
	
	if (S[0] != 1) {
		S = append([]int64{0}, S...)
	}
	if (S[len(S)-1] != N) {
		S = append(S, N)
	}
	ret += (S[0] - 1) / (K + 1)

	for i := 1; i < len(S); i++ {
		ret += (S[i] - S[i-1] - K - 1) / (K + 1)
	}

	return ret
}