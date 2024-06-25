package tunnel

import (
	"fmt"
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
	var time int64
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	for i := 0; i < int(K); i++ {
		fmt.Println("i=", i)
		if K <= 0 {
			return time
		}
		if B[i] - A[i] >= K {
			fmt.Printf("T > K | Current time: %d \t New time: %d\n", time, time + A[i] + K)
			time +=A[i] + K
			return time
		} else {
			fmt.Printf("ELSE | Current K: %d \t New K: %d\n", K, K - (B[i] - A[i]))
			K -= B[i] - A[i]
			if i == len(B) - 1 {
				i = -1
				fmt.Printf("ELSE | Current time: %d \t New time: %d\n", time, time + C)
				time += C
			}
		}
	}

	return time
}

func Solve(C int64, N int32, A []int64, B []int64, K int64) int64 {
	return getSecondsElapsed(C, N, A, B, K)
}