package leetcode

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	return deepPreorderTraversal(root, []int{})
}

func preorderTraversal1(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

func deepPreorderTraversal(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = deepPreorderTraversal(root.Left, ans)
	ans = deepPreorderTraversal(root.Right, ans)

	return ans
}
