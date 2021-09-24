package leetcode

func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0

	cMax := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for left <= right {
		if height[left] < height[right] {
			sum := (right - left) * height[left]
			left++
			ans = cMax(sum, ans)
		} else {
			sum := (right - left) * height[right]
			right--
			ans = cMax(sum, ans)
		}
	}

	return ans
}
