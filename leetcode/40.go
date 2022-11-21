package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int)bool{return candidates[i] < candidates[j]})
	ans := dfsCombinationSum2(candidates, target,0, []int{}, [][]int{})

	return ans
}

func dfsCombinationSum2(candidates []int, target, idx int, combination []int, ans [][]int)[][]int{
	for ; idx<len(candidates); idx++{
		if candidates[idx] == target{
			newCombination := make([]int, len(combination))
			copy(newCombination, combination)
			newCombination = append(combination, candidates[idx])
			ans = append(ans, newCombination)
		}
		if candidates[idx] > target{
			break
		}
		if candidates[idx] < target{
			newTarget := target - candidates[idx]
			newCombination := make([]int, len(combination))
			copy(newCombination, combination)
			newCombination = append(newCombination, candidates[idx])
			ans = dfsCombinationSum2(candidates, newTarget, idx+1, newCombination, ans)
		}
	}

	return ans
}

