package leetcode

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	items := []int{}

	ans := traversalTree(root, items)

	fmt.Printf("%v", ans)
	length := len(ans)
	for i, v := range ans {
		if i < length-1 && v >= ans[i+1] {
			return false
		}
	}

	return true
}

func traversalTree(node *TreeNode, items []int) []int {
	if node == nil {
		return items
	}

	items = traversalTree(node.Left, items)
	items = append(items, node.Val)
	items = traversalTree(node.Right, items)

	return items
}
