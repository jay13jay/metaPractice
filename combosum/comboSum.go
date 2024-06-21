package combosum

import "sort"

func ComboSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) 
	return combinationSum(candidates, target)
}

// candidates = {2, 3, 6, 7}
// target = 7
func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}

	// Base cases
	if target == 0 {
		return [][]int{{}}
	}
	if target < 0 {
		return nil
	}

	// Iterate through candidates
	for i := 0; i < len(candidates); i++ {
		c := candidates[i]
		// Skip duplicates to avoid generating duplicate combinations
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		remain := target - c

		// Recursively find combinations for the remaining target
		res := combinationSum(candidates[i:], remain)

		// If res is not nil (meaning valid combinations were found)
		if res != nil {
			// Append current candidate `c` to each combination in `res`
			for _, comb := range res {
				// Create a new combination by appending `c` to `comb`
				newComb := append([]int{c}, comb...)
				ret = append(ret, newComb)
			}
		}
	}

	return ret
}