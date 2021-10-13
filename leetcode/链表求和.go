package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	borrow := 0
	ans := &ListNode{}
	node := ans

	k := 0
	for !(l1 == nil && l2 == nil) {
		now := 0
		if l1 != nil {
			now += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			now += l2.Val
			l2 = l2.Next
		}

		now = now + borrow
		borrow = 0

		if now >= 10 {
			borrow = now / 10
			now = now % 10
		}

		newNode := &ListNode{Val: now}
		if k == 0 {
			ans = newNode
			node = ans
		} else {
			node.Next = newNode
			node = node.Next
		}

		k++
	}

	if borrow != 0 {
		newNode := &ListNode{Val: borrow}
		node.Next = newNode
	}

	return ans
}
