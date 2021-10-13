package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	firsts := []int{}
	for i := 0; i < len(matrix); i++ {
		firsts = append(firsts, matrix[i][len(matrix[i])-1])
	}

	k := BinarySearch(len(firsts), func(i int) bool {
		return firsts[i] >= target
	})

	if k < len(firsts) && firsts[k] == target {
		return true
	}

	if k == len(firsts) {
		return false
	}

	ans := BinarySearch(len(matrix[k]), func(i int) bool {
		return matrix[k][i] >= target
	})

	if ans < len(matrix[k]) && matrix[k][ans] == target {
		return true
	}

	return false
}

func BinarySearch(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i
	return i
}
