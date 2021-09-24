package leetcode

func getLinkedLength(head *ListNode) int {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	return length
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := getLinkedLength(head)
	ans := []*ListNode{}

	part := length / k // 每部分的个数
	others := length % k

	if length <= k {
		part = 1
		others = 0
	}

	nowNode := head
	for ; k > 0; k-- {
		partHead := nowNode
		partNode := partHead
		for p := part; p > 1; p-- {
			if partNode != nil {
				partNode = partNode.Next
			}
		}

		if others > 0 {
			if partNode != nil {
				partNode = partNode.Next
			}
			others--
		}

		if partNode != nil {
			nowNode = partNode.Next
			partNode.Next = nil
		} else {
			nowNode = partNode
		}

		ans = append(ans, partHead)
	}

	return ans
}
