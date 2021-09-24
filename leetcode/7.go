package leetcode

import "math"

func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}

		ans = ans*10 + x%10
		x = x / 10
	}

	return ans
}
