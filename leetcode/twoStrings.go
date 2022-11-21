package leetcode

func TwoStrings(a []string, b[] string) []string {
	l := len(a)
	ans := make([]string, 0)
	for i:=0; i<l; i++{

		ans = append(ans, hasSubString(a[i], b[i]))
	}

	return ans
}

func hasSubString(s1,s2 string) string{
	map1 := map[byte]int{}
	map2 := map[byte]int{}

	for _, b := range []byte(s1){
		map1[b] = 1
	}

	for _, b := range []byte(s2){
		map2[b] = 1
	}

	for key, _ := range map1{
		if map2[key] == 1{
			return "YES"
		}
	}

	return "NO"
}
