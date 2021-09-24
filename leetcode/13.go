package leetcode

var symbols = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
	' ': 0,
}

func romanToInt(s string) int {
	length := len(s)
	last := byte(' ')
	ans := 0
	for i := length - 1; i >= 0; i-- {
		if symbols[s[i]] < symbols[last] {
			ans -= symbols[s[i]]
		} else {
			ans += symbols[s[i]]
		}

		last = s[i]
	}

	return ans
}
