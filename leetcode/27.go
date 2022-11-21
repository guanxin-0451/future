package leetcode

func removeElement(nums []int, val int) int {
	i, j, ans :=0,len(nums)-1, len(nums)

	for i=0; i<=j; i++{
		if nums[i] == val{
			ans --
			for ; j>i;j--{
				if nums[j] == val{
					ans -=1
				}else{
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	return ans
}
