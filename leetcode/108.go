package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	return binaryFind(nums, 0, len(nums)-1)
}

func binaryFind(nums []int, left, right int) *TreeNode {
	mid := (left + right) / 2
	if left > right {
		return nil
	}
	node := &TreeNode{Val: nums[mid]}
	if left == right {
		return node
	}

	node.Left = binaryFind(nums, left, mid-1)
	node.Right = binaryFind(nums, mid+1, right)

	return node
}
