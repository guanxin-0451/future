package leetcode

func jump(nums []int) int {

	i := 0
	ans := 0
	target := len(nums)-1
	for {
		nowMax := i
		jumpTo := i
		for j :=0; j<nums[i];j++{
			if i+ nums[i + j]>=target{
				if j >0{
					return ans + 1
				}
				return ans
			}
			if i+ nums[i + j] >= nowMax{
				nowMax = i + nums[i+j]
				jumpTo = i+j
			}
		}
		i = jumpTo
		ans +=1

	}

}

