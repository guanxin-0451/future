package leetcode

var solutions [][]string
var find [][]byte

func makeFindMap(n int) {
	for i := 0; i < n; i++ {
		s := []byte{}
		for j := 0; j < n; j++ {
			s = append(s, '.')
		}

		find = append(find, s)
	}
}

func solveNQueens(n int) [][]string {
	makeFindMap(n)
	backtrack(n, 0)
	return solutions
}

func backtrack(n, t int) { // n 几行，t 当前放置到了第几行
	if t == n {
		ans := []string{}

		for _, i := range find {
			ans = append(ans, string(i))
		}

		solutions = append(solutions, ans)
	}

	for i := 0; i < n; i++ {
		if check(t, i, n) {
			find[t][i] = 'Q'
			backtrack(n, t+1)
			find[t][i] = '.'
		}
	}
}

func check(i, j, n int) bool {
	for k := 0; k < n; k++ {
		if k != i && find[k][j] == 'Q' {
			return false
		}
	}

	// diagonal（对角线判断）
	for k := 1; k < n; k++ {
		// 主对角线
		if i-k >= 0 && j-k >= 0 && find[i-k][j-k] == 'Q' {
			return false
		}
		if i+k < n && j+k < n && find[i+k][j+k] == 'Q' {
			return false
		}
		// 副对角线
		if i-k >= 0 && j+k < n && find[i-k][j+k] == 'Q' {
			return false
		}
		if i+k < n && j-k >= 0 && find[i+k][j-k] == 'Q' {
			return false
		}
	}

	return true

}
