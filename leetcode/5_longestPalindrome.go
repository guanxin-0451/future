package leetcode

func longestPalindrome(s string) string {

	ans := ""
	for i, _ := range s {
		l1 := i
		r1 := i

		ans = maxString(ans, getPalindrome(s, l1, r1))
		l2 := i
		r2 := i + 1
		ans = maxString(ans, getPalindrome(s, l2, r2))

	}

	return ans
}

func maxString(s1, s2 string) string {
	if len(s1) < len(s2) {
		return s2
	}
	return s1
}
func getPalindrome(s string, l, r int) string {
	length := len(s)

	for l >= 0 && r < length {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}

	return s[l+1 : r]
}
