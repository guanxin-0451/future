package leetcode

func pathSum(root *TreeNode, targetSum int) int {
	return findTreeM(root, targetSum)
}

func findTreeM(node *TreeNode, target int) int {
	ans := 0
	if node == nil {
		return ans
	}

	ans += findTreeD(node, 0, 0, target)
	ans1 := findTreeM(node.Left, target)
	ans2 := findTreeM(node.Right, target)
	return ans + ans1 + ans2

}
func findTreeD(node *TreeNode, num int, ans int, targetSum int) int {
	if node == nil {
		return ans
	}

	num = num + node.Val
	// fmt.Printf("\n%v %v %v\n", node.Val, num, ans)
	if num == targetSum {
		ans += 1
	}

	ans1 := findTreeD(node.Left, num, ans, targetSum)
	ans2 := findTreeD(node.Right, num, ans, targetSum)

	ans = ans1 + ans2

	return ans
}

func pathSum1(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}
	dfs(root, 0)
	return
}
