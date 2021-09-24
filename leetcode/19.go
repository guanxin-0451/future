package leetcode

// 删除倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, head
	for i := 0; first != nil; i++ {
		if i < n {
			first = first.Next
			continue
		}

		second = second.Next
		first = first.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
