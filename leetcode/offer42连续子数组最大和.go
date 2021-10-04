package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		m = max(nums[i], m)
	}

	return m
}
