package leetcode

// 最佳观光组合
func maxScoreSightseeingPair(values []int) int {

	maxVWitght := 0
	ans := 0
	for i, v := range values {

		if i > 0 {
			ans = max(maxVWitght+v-i, ans)
		}
		maxVWitght = max(v+i, maxVWitght)
	}

	return ans

}
