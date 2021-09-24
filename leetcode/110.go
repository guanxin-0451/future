package leetcode

// 3:28 -3:44

func isBalanced(root *TreeNode) bool {
	if validDistance(root) >= 10000 {
		return false
	}

	return true
}

func validDistance(node *TreeNode) (height int) {
	if node == nil {
		return 0
	}

	leftHeight := validDistance(node.Left)
	rightHeight := validDistance(node.Right)
	if (leftHeight-rightHeight) < -1 || (leftHeight-rightHeight) > 1 {
		return 10000
	}

	return max(leftHeight, rightHeight)
}
