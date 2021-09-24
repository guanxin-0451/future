package leetcode

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	return subPermuteUnique(0, nums, ans)
}

func subPermuteUnique(iter int, nums []int, ans [][]int) [][]int {
	lenNums := len(nums)
	dict := make(map[int]int)

	for i := iter; i < lenNums; i++ {
		newNums := make([]int, lenNums)
		copy(newNums, nums)
		if dict[newNums[i]] == 1 {
			continue
		}

		dict[newNums[i]] = 1

		newNums[iter], newNums[i] = newNums[i], newNums[iter]
		if iter == lenNums-1 {
			ans = append(ans, newNums)
			fmt.Printf("\nans:%v", ans)
			continue
		}

		ans = subPermuteUnique(iter+1, newNums, ans)
	}

	return ans
}
