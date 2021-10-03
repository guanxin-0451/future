//
//func main() {
//	////leetcode.Permute([]int{1, 2,3})
//	//node1 := leetcode.TreeNode{Val: 1}
//	//node1.Left = &leetcode.TreeNode{Val: 2}
//	//node1.Right = &leetcode.TreeNode{Val: 3}
//	//node1.Left.Left = &leetcode.TreeNode{Val: 4}
//	//node1.Left.Right = &leetcode.TreeNode{Val: 5}
//	//node1.Left.Left.Left = &leetcode.TreeNode{Val: 7}
//	//node1.Left.Left.Left.Right = &leetcode.TreeNode{Val: 8}
//	//
//	//ans := leetcode.PostorderTraversal(&node1)
//	//fmt.Printf("\n%v\n", ans)
//}

package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{1, 1, 1}
	b := []int{2, 2, 2}
	c := []int{3, 3, 3}

	ans := calculate(a, b, c)
	fmt.Printf("\n%v\n", ans)
}

func calculate(in ...[]int) []int {
	result := make([]int, 3)
	var wg sync.WaitGroup

	wg.Add(len(in))
	for i, s := range in {
		s1 := s
		i1 := i
		go func() {
			defer wg.Add(-1)
			sum := 0
			for _, value := range s1 {
				sum += value
			}
			result[i1] = sum
		}()
	}

	wg.Wait()

	return result
}
