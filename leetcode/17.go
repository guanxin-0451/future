package leetcode

var m = map[rune][]rune{
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'v', 't', 'u'},
	'9': []rune{'w', 'x', 'y','z'},


}

func letterCombinations(digits string) []string {
	var ans []string
	for _, c := range digits{
		var newAns []string
		for _, l := range m[c]{
			if len(ans) == 0{
				newAns = append(newAns, string(l))
			}else{
				for _, str := range ans{
					newAns = append(newAns, str+string(l))
				}
			}
		}
		ans = newAns
	}

	return ans
}
