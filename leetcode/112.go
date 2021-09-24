package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return countNodeWeight(root, 0, targetSum, false)
}

func countNodeWeight(node *TreeNode, weight, targetSum int, has bool) bool {
	if node.Left == nil && node.Right == nil {
		return weight+node.Val == targetSum || has
	}

	leftHas := false
	rightHas := false

	if node.Left != nil {
		leftHas = countNodeWeight(node.Left, weight+node.Val, targetSum, has)
	}
	if node.Right != nil {
		rightHas = countNodeWeight(node.Right, weight+node.Val, targetSum, has)

	}

	return leftHas || rightHas
}
