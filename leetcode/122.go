package leetcode

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	ans := 0

	for i, v := range prices {

		if i > 0 && v-prices[i-1] > 0 {
			ans += v - prices[i-1]
		}
	}

	return ans
}
