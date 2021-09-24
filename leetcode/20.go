package leetcode

var leftMap = map[int32]int32{
	'[': ']',
	'{': '}',
	'(': ')',
}

func isLeft(c int32) bool {
	_, ok := leftMap[c]
	return ok
}

func isValid(s string) bool {
	stack := []int32{}
	for _, c := range s {
		if isLeft(c) {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if leftMap[stack[len(stack)-1]] != c {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
