package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	// write code here
	var queue []*TreeNode
	if root == nil {
		return nil
	}
	queue = append(queue, root)

	var ans [][]int

	for level := 1; len(queue) != 0; level++ {
		var newQueue []*TreeNode
		len := len(queue)
		row := make([]int, len)

		for i, n := range queue {
			if n.Left != nil {
				newQueue = append(newQueue, n.Left)
			}

			if n.Right != nil {
				newQueue = append(newQueue, n.Right)
			}

			k := i

			if level%2 == 0 {
				k = len - 1 - i
			}

			row[k] = n.Val
		}

		ans = append(ans, row)
		queue = newQueue
	}

	return ans
}
