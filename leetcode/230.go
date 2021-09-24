package leetcode

func kthSmallest(root *TreeNode, k int) int {
	ans := findTree(root, []int{})
	if k > len(ans) || k < 0 {
		return -1
	}

	return ans[k-1]
}

func findTree(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}

	ans = findTree(root.Left, ans)
	ans = append(ans, root.Val)
	ans = findTree(root.Right, ans)
	return ans
}
