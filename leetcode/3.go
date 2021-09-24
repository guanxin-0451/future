package leetcode

// 无重复字符的最长子串

func LengthOfLongestSubstring(s string) int {
	var subLen, r, ans int
	l := len(s)
	dict := make(map[byte]int)

	for i, _ := range s {
		for r < l && dict[s[r]] == 0 {
			//fmt.Printf("r:%d  s[r]: %v\n", r, string(s[r]))
			dict[s[r]]++
			r++
			subLen++
		}

		if subLen > ans {
			ans = subLen
		}

		dict[s[i]] = 0
		subLen -= 1
	}

	return ans
}
