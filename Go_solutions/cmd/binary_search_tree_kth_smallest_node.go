package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    for {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        stack, root = stack[:len(stack)-1], stack[len(stack)-1]
        k--
        if k == 0 {
            return root.Val
        }
        root = root.Right
    }
}
 */


func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack = stack[:len(stack)-1]
		root = stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
