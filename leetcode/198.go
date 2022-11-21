package leetcode
func rob(nums []int) int {
	//numsMax := make([]int, len(nums))
	numsLeft, numsRight := 0, 0

	for i, num := range nums{
		if i == 0{
			numsLeft = num
			numsRight = num
		}else if i == 1{
			numsRight = max(num, numsLeft)
		} else{
			numsMAx := max(num+numsLeft, numsRight)
			numsLeft = numsRight
			numsRight = numsMAx
		}

	}

	return numsRight
}


func max(x, y int) int{
	if x > y{return x}
	return y
}