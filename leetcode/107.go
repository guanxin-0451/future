package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	ans := [][]int{}

	if root == nil {
		return nil
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		row := []int{}
		for i := 0; i < length; i++ {
			head := queue[0]
			queue = queue[1:]
			row = append(row, head.Val)

			if head.Left != nil {
				queue = append(queue, head.Left)
			}

			if head.Right != nil {
				queue = append(queue, head.Right)
			}

		}

		ans = append(ans, row)

	}
	ans1 := [][]int{}
	for i := len(ans) - 1; i >= 0; i-- {
		ans1 = append(ans1, ans[i])
	}

	return ans1
}
