package leetcode

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {

	if left == right {
		if target <= nums[right] {
			return right
		} else {
			return right + 1
		}
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}

	if target > nums[mid] {
		if mid+1 > right {
			return right + 1
		}
		return binarySearch(nums, target, mid+1, right)
	} else {
		if mid-1 < left {
			return left
		}
		return binarySearch(nums, target, left, mid-1)

	}
}
