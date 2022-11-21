package leetcode

func isPresent(node *TreeNode, val int)int{

	for node != nil{
		if node.Val == val{
			return 1
		}

		if node.Val < val{
			node = node.Left
		}else {
			node = node.Right
		}
	}

	return 0
}