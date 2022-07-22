package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	listArr := make([]*ListNode, 0)
	for head.Next != nil {
		listArr = append(listArr, head)
		head = head.Next
	}
	listArr = append(listArr, head)

	for i := len(listArr)-1; i > 0; i-- {
		listArr[i].Next = listArr[i-1]
	}
	listArr[0].Next = nil
	return listArr[len(listArr)-1]
}
