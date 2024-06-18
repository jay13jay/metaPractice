package combosum

func ComboSum(candidates []int, target int) [][]int { 
	return combinationSum(candidates, target)
}

// ret := combosum.ComboSum([]int{2, 3, 6, 7}, 7)

// maps checked values to what will add up to sum
var iMap = make(map[int][]int) 
var cMap = make(map[int]bool)
// candidates = {2, 3, 6, 7}
// target = 7
func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if target == 0 {
		return ret
	}
	if target < 0 {
		return nil
	}

	for i := 0; i < len(candidates); i++ {
		c := candidates[i]
		remain := target - c
		res := combinationSum(candidates, remain)
		if res != nil{
			return ()
		}
		return nil
	}

	return ret
}