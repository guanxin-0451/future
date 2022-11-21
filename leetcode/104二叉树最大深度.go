package leetcode

func maxDepth(root *TreeNode) int {
	return traversalTree1(root, 0)
}

func traversalTree1(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	depth++

	return max(traversalTree1(node.Left, depth), traversalTree1(node.Right, depth))
}
