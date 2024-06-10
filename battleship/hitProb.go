package battleship

func GetHitProbability(R int32, C int32, G [][]int32) float64 {
  // Write your code here
  var sum int
  for _, v := range G {
			for _, val := range v {
				if val == 1 {
					sum++
				}
			}
		}

		area := R * C
		prob := float64(sum) / float64(area)
  return float64(prob)
}

