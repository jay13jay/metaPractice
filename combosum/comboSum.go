package combosum

import "fmt"

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
	for i := 0; i < len(candidates); i++ {
		var tempSum int = 0
		tempSum += candidates[i]
		tempRet := []int{candidates[i]}
		comp := target - tempSum
		// fmt.Printf("MAIN tempSum: %d\ttempRet: %v\tcomp: %d\n", tempSum, tempRet, comp)

		for comp > 0 {
			fmt.Printf("tempSum: %d\ttempRet: %v\tcomp: %d\n", tempSum, tempRet, comp)

			if cMap[tempSum] {
				fmt.Printf("cmap[tempSum] triggered\n")
				tempRet = append(tempRet, iMap[tempSum]...)
				ret = append(ret, tempRet)
				fmt.Printf("ret: %v\n", ret)
			}
			if cMap[comp] {
				iMap[comp] = append(iMap[comp], tempRet...)
			}
			// iMap[comp] = append(iMap[comp], tempRet...)
			cMap[comp] = true

			tempSum += candidates[i]
			tempRet = append(tempRet, candidates[i])
			comp = target - tempSum
		}

		if comp == 0 {
			ret = append(ret, tempRet)
		}
	}
	// sort.IntSlice(ret)
	return ret
}