package leetcode

// 三角形路径和

func getMinFather(triangle [][]int, x, y int) int {
	fLen := len(triangle[x-1])

	if y == 0 {
		return triangle[x-1][0]
	}

	if y == fLen {
		return triangle[x-1][fLen-1]
	}

	return min(triangle[x-1][y], triangle[x-1][y-1])
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	l := len(triangle)
	ans := 0
	for i := 1; i < l; i++ {
		if i == l-1 {
			ans = triangle[i][0] + getMinFather(triangle, i, 0)
		}
		for j, value := range triangle[i] {
			triangle[i][j] = value + getMinFather(triangle, i, j)
			if i == l-1 {
				ans = min(ans, triangle[i][j])
			}
		}
	}

	return ans

}
