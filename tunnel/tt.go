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

	// return value: total time needed to reack K time in tunnel
	var ret int64 = 0
	var tnl map[int]int64 = make(map[int]int64)
	var lapTotal int64 = 0
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	// get length of tunnes
	for i := 0; i < int(N); i++ {
		tnl[i] = B[i] - A[i]
		lapTotal += tnl[i]
	}
	fmt.Printf("Tunnels: %v\tTotal time per lap: %d \t tTime needed: %d\n", tnl, lapTotal, K)
	laps := K / lapTotal
	remain := K % lapTotal
	if K > lapTotal {
		fmt.Printf("K: %d \t laptotal: %d\n", K, lapTotal)
		ret += C * laps
		// fmt.Println("Laps -1 ", laps - 1)
	}
	if remain == 0 {
		ret += B[len(B)-1]
		return ret
	}

	fmt.Printf("Laps needed: %d\t remain: %d\tcurrent return: %d\n", laps, remain, ret)

	for i := 0; i < int(N); i++ {
		fmt.Printf("toadd: %d remain: %d return: %d\n", (B[i] - A[i]), remain, ret)
		if (B[i] - A[i] < remain) {
			remain -= B[i] - A[i]
			// fmt.Println("remain: ", remain)
			// ret += B[i]
			continue
		} else if remain > 0 {
			fmt.Printf("Last add: %d remain: %d\n", A[i] + remain, remain)
			ret += A[i] + remain
			break
		}
	}
	fmt.Printf("Ret: %d\n", ret)
  return ret
}

func Solve(C int64, N int32, A []int64, B []int64, K int64) int64 {
	return getSecondsElapsed(C, N, A, B, K)
}