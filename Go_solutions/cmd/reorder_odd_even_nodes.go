package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || head.Next.Next == nil {
		return head
	}
	headStored := head
	for head.Next.Next != nil {
		head = deal3nodes(head)
	}
	return headStored
}

func deal3nodes(head *ListNode) *ListNode {
	secondNode := head.Next
	thirdNode := head.Next.Next
	head.Next = thirdNode
	thirdNode.Next = secondNode
	return thirdNode
}
