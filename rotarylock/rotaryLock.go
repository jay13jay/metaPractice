package rotarylock

// N = number of digits on lock
// M = number of digits in code
// C = code to unlock

// 1, 2, 3, 4, 5

func getMinCodeEntryTime(N int32, M int32, C []int32) int64 {
  // Write your code here
	current := int32(1)
	var time int64 = 0
	for i := int32(0); i < M; i++ {
		if C[i] == current {
			continue
		}
		var back int32
		var forward int32
		if C[i] < current {
			back = current - C[i]
			forward = N - current + C[i]
		} else {
			back = N - C[i] + current
			forward = C[i] - current
		}
		if back < forward {
			time += int64(back)
		} else {
			time += int64(forward)
		}
		current = C[i]
	}
  return time
}

func Solve(N, M int32, C []int32) int64 {
	return getMinCodeEntryTime(N, M, C)
}