package leetcode

import "container/list"

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func bfs(grid [][]byte, local []int) [][]byte {

	xLen, yLen := len(grid), len(grid[0])
	queue := list.New()

	queue.PushBack(local)
	grid[local[0]][local[1]] = 0

	for queue.Len() > 0 {
		pop := queue.Front()
		x, y := pop.Value.([]int)[0], pop.Value.([]int)[1]

		for _, d := range directions {
			if x+d[0] >= 0 && x+d[0] < xLen && y+d[1] >= 0 && y+d[1] < yLen && grid[x+d[0]][y+d[1]] == '1' {
				queue.PushBack([]int{x + d[0], y + d[1]})
				grid[x+d[0]][y+d[1]] = '0'
			}
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
