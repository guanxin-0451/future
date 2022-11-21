package leetcode

func rotate(matrix [][]int)  {
	length := len(matrix) - 1
	l, r := 0, len(matrix) - 1
	for l<r{
		for x:=l;x<r;x++{
			// 本行坐标 x, l - r
			y := length-r
			k := matrix[y][x]
			matrix[y][x] = matrix[length-x][y]
			matrix[length-x][y] = matrix[length-y][length-x]
			matrix[length-y][length-x] = matrix[x][length-y]
			matrix[x][length-y] = k
		}
		l++
		r--
	}

}
