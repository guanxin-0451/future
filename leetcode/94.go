package leetcode

func inorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		vals = append(vals, stack[len(stack)-1].Val)
		stack = stack[:len(stack)-1]
	}
	return
}
