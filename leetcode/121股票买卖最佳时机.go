package leetcode

// 维护截止当前的最小价格，计算当前的价格和最小价格的价格差
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	maxCount := 0

	for _, v := range prices[1:] {
		if v < minPrice {
			minPrice = v
		} else {
			maxCount = max(maxCount, v-minPrice)
		}
	}

	return maxCount
}
