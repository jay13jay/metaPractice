package scoreboard

import (
	"fmt"
	"sort"
)

func getMinProblemCount(N int32, S []int32) int32 {
  // Write your code here
	var scores int32
	var odd, even bool
	max := int32(0)

	// Sort slices as we only need the maximum value
	sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })
	max = S[len(S)-1]

	// Check if there are any odd/even numbers in the slice

	for _, v := range S {
		if v % 2 == 0 {
			even = true
			if odd {
				break
			}
		} else {
			odd = true
			if even {
				break
			}
		}
	}

	if odd && even {
		fmt.Println("Current max: ", max)
		scores += 2
		max -= 2
	} else if odd {
			scores += 1
			max -= 1
	} 
	scores += max / 2

  return scores
}


func Solve(N int32, S []int32) int32 {
	return getMinProblemCount(N, S)
}