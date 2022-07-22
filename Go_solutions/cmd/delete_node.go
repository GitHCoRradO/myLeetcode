package main

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	// nextNode will never be nil
	nextNode := node.Next
	thirdNode := nextNode.Next
	if thirdNode == nil {
		node.Val = nextNode.Val
		node.Next = nil
		return
	} else {
		node.Val = nextNode.Val
		node.Next = thirdNode
		return
	}
}
