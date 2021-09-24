package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	nilHead := &ListNode{0, nil}
	nilHead.Next = head

	if head == nil || head.Next == nil {
		return head
	}

	last := head

	for node := head.Next; node != nil; node = node.Next {
		if node.Val == last.Val {
			last.Next = node.Next
		} else {
			last = last.Next
		}
	}

	return nilHead.Next
}
