package leetcode

func swapPairs(head *ListNode) *ListNode {
	nilHead := &ListNode{}
	nilHead.Next = head
	node := nilHead


	for {
		whiteNode := node.Next
		if whiteNode != nil && whiteNode.Next != nil{
			blackNode := whiteNode.Next
			whiteNode.Next = blackNode.Next
			blackNode.Next = whiteNode
			node.Next = blackNode
			node = whiteNode
		}else{
			break
		}
	}

	return nilHead.Next
}