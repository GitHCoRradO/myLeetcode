package main

import "sort"

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	vals := make([]int, 0)
	for head.Next != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	vals = append(vals, head.Val)
	sort.Ints(vals)
	newNodes := make([]*ListNode, 0)
	for _, val := range vals {
		newNodes = append(newNodes, &ListNode{Val: val})
	}
	for i := 0; i < len(newNodes) - 1; i++ {
		newNodes[i].Next = newNodes[i+1]
	}
	return newNodes[0]
}
