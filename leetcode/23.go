package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if r-l <= 1 {
		if r-l < 0 {
			return nil
		}
		if r-l == 0 {
			return lists[l]
		}

		return mergeTowList(lists[l], lists[r])
	}

	ans1 := merge(lists, l, (l+r)/2)
	ans2 := merge(lists, (l+r)/2+1, r)

	return mergeTowList(ans1, ans2)

}
func mergeTowList(list1, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		}

		return list1
	}

	head := ListNode{}
	node1, node2 := list1, list2
	now := &head

	for {
		if node1 == nil || node2 == nil {
			if node1 != nil {
				now.Next = node1
				now = now.Next
				node1 = node1.Next
				continue
			}

			if node2 != nil {
				now.Next = node2
				now = now.Next
				node2 = node2.Next
				continue
			}

			break
		}

		if node1.Val <= node2.Val {
			now.Next = node1
			now = now.Next
			node1 = node1.Next
		} else {
			now.Next = node2
			now = now.Next
			node2 = node2.Next
		}
	}

	return head.Next
}
