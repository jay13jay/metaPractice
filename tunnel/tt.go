package tunnel

import (
	"sort"
)

/*
 C total len of track 			| 10 [0-9]
 N # of tunnels 						| 2
 A tunnel start points			| [1, 6]
 B tunnel end points				| [3, 7]
 K time in tunnel to reach	| 15
*/
func getSecondsElapsed(C int64, N int32, A []int64, B []int64, K int64) int64 {
  // Write your code here
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	var tunnelTime int64
	for i := 0; i < int(N); i++ {
		tunnelTime += B[i] - A[i]
	}
	laps := K / tunnelTime
	remain := K % tunnelTime

	if remain == 0 {
		return C * laps - (C - B[N-1])
	}

	i := 0
	for remain > 0 {
		remain -= B[i] - A[i]
		i++
	}
	return C * laps + B[i - 1] + remain
}

func Solve(C int64, N int32, A []int64, B []int64, K int64) int64 {
	return getSecondsElapsed(C, N, A, B, K)
}