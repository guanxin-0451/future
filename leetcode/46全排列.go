package leetcode

import "fmt"

func Permute(nums []int) [][]int {
	return permute(nums)
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	return subPermute(0, nums, ans)
}

func subPermute(iter int, nums []int, ans [][]int) [][]int {
	lenNums := len(nums)
	for i := iter; i < lenNums; i++ {
		newNums := make([]int, lenNums)
		copy(newNums, nums)
		newNums[iter], newNums[i] = newNums[i], newNums[iter]
		if iter == lenNums-1 {
			ans = append(ans, newNums)
			fmt.Printf("\nans:%v", ans)
			continue
		}

		ans = subPermute(iter+1, newNums, ans)
	}

	return ans
}
