package main

import (
	"fmt"
	"future/leetcode"
)

func main() {
	//leetcode.Permute([]int{1, 2,3})
	node1 := leetcode.TreeNode{Val: 1}
	node1.Left = &leetcode.TreeNode{Val: 2}
	node1.Right = &leetcode.TreeNode{Val: 3}
	node1.Left.Left = &leetcode.TreeNode{Val: 4}
	node1.Left.Right = &leetcode.TreeNode{Val: 5}
	node1.Left.Left.Left = &leetcode.TreeNode{Val: 7}
	node1.Left.Left.Left.Right = &leetcode.TreeNode{Val: 8}

	ans := leetcode.PostorderTraversal(&node1)
	fmt.Printf("\n%v\n", ans)
}
