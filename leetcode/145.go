package leetcode

func PostorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

func PostorderTraversal2(root *TreeNode) (vals []int) {
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
