package leetcode

import "math"

// 深度优先搜索+记忆化搜索过的路径
func coinChange(coins []int, amount int) int {
	memory := make(map[int]int)
	for _, i := range coins{
		memory[i] = 1
	}
	var dfs func(amount int)(less int)
	dfs = func (amount int) (less int) {
		if amount < 0{
			return - 1
		}
		if amount == 0{
			return 0
		}

		if _, ok := memory[amount]; ok{
			return memory[amount]
		}

		can := false
		minStep := math.MaxInt32
		for _, i := range coins {
			if amount - i >= 0 {
				if amount -1 == 0{
					can = true
					minStep = 1
				}else{
					step := dfs(amount - i)
					if step >= 0{
						can = true
						minStep = min(minStep, 1+step)
					}
				}
			}
		}

		if can {
			memory[amount] = minStep
			return minStep
		}

		memory[amount] = -1
		return -1
	}

	return dfs(amount)
}

func min(a, b int)int{
	if a < b{
		return a
	}

	return b
}

