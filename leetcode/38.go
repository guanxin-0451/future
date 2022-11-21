package leetcode

var sList = GetList()
func countAndSay(n int) string {
	return sList[n]
}

func GetList() []string{
	ans := make([]string, 31)
	ans[1] = "1"
	for i:=2; i<= 30; i++{
		newStr := ""
		num := 1
		oldB := '0'
		//fmt.Printf("str :%v\n", ans[i-1])
		for _, b := range ans[i-1]{
			//fmt.Printf("b: %v %v\n", b, string(b))
			if b == oldB{
				num++
			}else{
				if oldB != '0'{
					//fmt.Printf("%v %v, %v, %v\n",ans[i-1], num, oldB, string(oldB))
					newStr = newStr + string(int32(num)+'1'-int32(1)) + string(oldB)
				}
				oldB = b
				num = 1
			}
		}

		newStr = newStr + string(int32(num)+'1'-int32(1)) + string(oldB)

		ans[i] = newStr
	}

	return ans
}

