package leetcode

import "container/list"

func bfs(grid [][]byte, local []int) [][]byte {
	xLen, yLen := len(grid), len(grid[0])
	queue := list.New()

	queue.PushBack(local)
	grid[local[0]][local[1]] = 0

	for queue.Len() > 0 {
		pop := queue.Front()
		x, y := pop.Value.([]int)[0], pop.Value.([]int)[1]
		if x-1 >= 0 && grid[x-1][y] == '1' {
			queue.PushBack([]int{x - 1, y})
			grid[x-1][y] = 0
		}
		if x+1 < xLen && grid[x+1][y] == '1' {
			queue.PushBack([]int{x + 1, y})
			grid[x+1][y] = 0

		}
		if y-1 >= 0 && grid[x][y-1] == '1' {
			queue.PushBack([]int{x, y - 1})
			grid[x][y-1] = 0

		}
		if y+1 < yLen && grid[x][y+1] == '1' {
			queue.PushBack([]int{x, y + 1})
			grid[x][y+1] = 0
		}

		queue.Remove(pop)
	}

	return grid
}
func numIslands(grid [][]byte) int {
	ans := 0

	for rowIter, row := range grid {
		for lineIter, _ := range row {
			if grid[rowIter][lineIter] == '1' {
				ans += 1
				grid = bfs(grid, []int{rowIter, lineIter})
			}
		}
	}

	return ans
}
