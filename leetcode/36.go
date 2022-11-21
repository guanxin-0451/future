package leetcode

func isValidSudoku(board [][]byte) bool {
	nineMap := make(map[int]map[byte]int)
	rowMap := make(map[int]map[byte]int)
	colMap := make(map[int]map[byte]int)

	for i := 0; i <= 8; i++ {
		nineMap[i] = make(map[byte]int)
		rowMap[i] = make(map[byte]int)
		colMap[i] = make(map[byte]int)
	}

	for y, rows := range board {
		for x, value := range rows {
			if value == '.' {
				continue
			}
			nineKey := x/3 + y/3*3
			nineMap[nineKey][value] += 1
			rowMap[x][value] += 1
			colMap[y][value] += 1
			if nineMap[nineKey][value] > 1 || rowMap[x][value] > 1 || colMap[y][value] > 1 {
				return false
			}
		}
	}

	return true
}
