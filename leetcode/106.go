package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}

	// fmt.Printf("\n%v %v %v %v\n", inorder, postorder , i, postorder[len(postorder)-1])

	// inorder[:i] 左子树长度
	root.Left = buildTree(inorder[:i], postorder[0:len(inorder[:i])])
	if i+1 > len(inorder)-1 {
		root.Right = nil
	} else {
		root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
	}
	return root
}
