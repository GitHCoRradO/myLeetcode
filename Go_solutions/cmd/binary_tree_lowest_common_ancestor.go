package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, q, p)
	right := lowestCommonAncestor(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
