package leetcode

// 逐个计算 超时
func findIntegersF(n int) int {
	ans := n + 1
	for i := 0; i <= n; i++ {
		lastFlag := false
		for j := i; j > 0; j = j >> 1 {
			if j&1 == 1 {
				if lastFlag {
					ans--
					break
				}
			}

			lastFlag = j&1 == 1
		}
	}

	return ans
}

func findIntegers(n int) (ans int) {
	dp := [31]int{1, 1}
	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	for i, pre := 29, 0; i >= 0; i-- {
		val := 1 << uint(i)
		if n&val > 0 {
			ans += dp[i+1]
			if pre == 1 {
				break
			}
			pre = 1
		} else {
			pre = 0
		}
		if i == 0 {
			ans++
		}
	}
	return
}
