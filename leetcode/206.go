package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 翻转链表：核心就是新增1个空的节点，不停的让从第一个节点开始反向指向他

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return prev
}
