package leetcode

func missingNumber(nums []int) int {
	count := 0
	for _, i := range nums {
		count += i
	}

	ans := (1+len(nums))*len(nums)/2 - count

	return ans
}
