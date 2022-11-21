package leetcode

func countDuplicates(nums []int)int{
	m := make(map[int]int)
	ans := 0

	for _, num := range nums{
		m[num] = m[num]+1
		if m[num] == 2{
			ans ++
		}
	}

	return ans
}
