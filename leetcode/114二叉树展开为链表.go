package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	nodes := traversalTree2(root, []*TreeNode{})
	length := len(nodes)

	for i := 0; i < length-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}

	root = nodes[0]
}

func traversalTree2(node *TreeNode, ans []*TreeNode) []*TreeNode {
	if node == nil {
		return ans
	}

	ans = append(ans, node)
	ans = traversalTree2(node.Left, ans)
	ans = traversalTree2(node.Right, ans)

	return ans
}
