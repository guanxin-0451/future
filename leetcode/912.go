package leetcode

func sortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) []int {
	var mid int
	if l >= r {
		return nums

	}

	nums, mid = QuickSortPart(nums, l, r)

	nums = quickSort(nums, l, mid-1)
	nums = quickSort(nums, mid+1, r)

	return nums
}

func QuickSortPart(nums []int, left, right int) ([]int, int) {
	if right <= left {
		return nums, left
	}

	tmp := nums[left]
	local := left
	l, r := left, right

	for l < r {
		for l <= r && tmp <= nums[r] {
			r--
		}

		if l > r {
			break
		}

		nums[l] = nums[r]
		nums[r] = tmp
		local = r

		for l <= r && tmp >= nums[l] {
			l++
		}

		if l > r {
			break
		}

		nums[r] = nums[l]
		nums[l] = tmp
		local = l
	}

	return nums, local

}
