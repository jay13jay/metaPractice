package hops

/*
N: 6 - 			# of pads
F: 3 - 			# of frogs
P: 5, 2, 4 	# occuppied pads
Pads:
1: false
2: true
3: false
4: true
5: true
6: false
*/
func getSecondsRequired(N int64, F int32, P []int64) int64 {
  // Write your code here
	frog1 := P[0]
	for i := 1; i < int(F); i++ {
		if P[i] < frog1 {
			frog1 = P[i]
		}
	}

  return N - frog1
}

func SolveHops(N int64, F int32, P []int64) int64 {
  return getSecondsRequired(N, F, P)
}