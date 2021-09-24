package leetcode

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		m := make(map[int]int)
		for _, q := range points {
			m[(p[0]-q[0])*(p[0]-q[0])+(p[1]-q[1])*(p[1]-q[1])] += 1
		}

		for _, v := range m {
			ans += v * (v - 1)
		}
	}
	return ans

}
