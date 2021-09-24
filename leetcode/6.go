package leetcode

// N字形变换
func convert(s string, numRows int) string {
	ans := ""
	length := len(s)
	maxDistance := (numRows - 1) * 2

	if numRows == 1 {
		maxDistance = 1
	}

	for i := 0; i < numRows; i++ {
		j := i
		distanceLeft, distanceRight := maxDistance-i*2, i*2
		flag := true //左距离还是右距离

		for j < length {
			if flag {
				if distanceLeft != 0 {
					ans += string(s[j])
					j = j + distanceLeft
				}
			} else {
				if distanceRight != 0 {
					ans += string(s[j])
					j = j + distanceRight
				}
			}

			flag = !flag
		}
	}

	return ans
}
