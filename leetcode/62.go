package leetcode

func uniquePaths(m int, n int) int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, m+1)

		dp[i][1] = 1
	}
	for i := 1; i <= m; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n][m]
}
