package leetcode

import "sort"

var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	Dfs([]int{}, nums, 0)
	return res
}
func Dfs(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		Dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}
