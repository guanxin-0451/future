package leetcode

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	leftMax := 0
	rightMax := 0
	left := 0
	right := len(height) - 1
	ans := 0

	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
