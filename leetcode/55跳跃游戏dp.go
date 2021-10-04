package leetcode

// 计算每个下标位置的可以到达的最远距离：nums[i] = max(nums[i-1] - 1, nums[i])
func canJump(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == 0 { // 说明到当前的上一个节点已经不能往下跳了
			return false
		}

		nums[i] = max(nums[i-1]-1, nums[i])
	}

	return true
}
