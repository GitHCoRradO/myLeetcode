package main

type listnode struct {
	Val int
	Next *listnode
}

var visitedListNodes map[*listnode]bool

func hasCycle(head *listnode) bool {
	if head.Next == nil {
		return false
	}
	visitedListNodes = map[*listnode]bool{}
	visitedListNodes[head] = true
	for head.Next != nil {
		if _, exists := visitedListNodes[head.Next]; !exists {
			visitedListNodes[head.Next] = true
			head = head.Next
		} else {
			return true
		}
	}
	return false
}
