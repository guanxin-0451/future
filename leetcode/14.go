package leetcode

func longestCommonPrefix(strs []string) string {
	ans := ""
	i := 0
	for {
		if i >= len(strs[0]) {
			return ans
		}

		char := strs[0][i]
		for _, s := range strs {
			if i >= len(s) {
				return ans
			}
			if s[i] != char {
				return ans
			}
		}

		i += 1
		ans += string(char)
	}

	return ans
}
