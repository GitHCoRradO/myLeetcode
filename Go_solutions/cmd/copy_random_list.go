package main

type node struct {
	Val int
	Next *node
	Random *node
}

var visited map[*node]*node

func copyRandomList(head *node) *node {
	if head == nil {
		return nil
	}
	visited = make(map[*node]*node)
	var oldNode = head
	var newNode = &node{Val: oldNode.Val}
	visited[oldNode]= newNode
	for oldNode != nil {
		newNode.Next = getClonedNode(oldNode.Next)
		newNode.Random = getClonedNode(oldNode.Random)
		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return visited[head]
}

func getClonedNode(n *node) *node {
	if n == nil {
		return nil
	} else {
		if v, exists := visited[n]; exists {
			return v
		} else {
			newNode := node{Val: n.Val}
			visited[n] = &newNode
			return visited[n]
		}
	}
}


