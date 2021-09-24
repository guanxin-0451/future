package leetcode

func superEggDrop(k int, n int) int {
	ans := 0
	for k > 1 {
		n = (n + 1) / 2
		ans += 1
		k--
	}

	ans += n - 1
	return ans
}
