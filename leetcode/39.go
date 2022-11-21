package leetcode


import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int)bool{return candidates[i] < candidates[j]})
	return dfsCombinationSum(candidates, target,0, []int{}, [][]int{})
}

func dfsCombinationSum(candidates []int, target, idx int, combination []int, ans [][]int)[][]int{
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
			ans = dfsCombinationSum(candidates, newTarget, idx, newCombination, ans)
		}
	}
	go combinationSum();
	return ans
}

