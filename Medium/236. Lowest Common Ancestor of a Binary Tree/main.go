package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == q || root == p {
		return root
	}
	r, l := lowestCommonAncestor(root.Right, p, q), lowestCommonAncestor(root.Left, p, q)
	if r != nil && l != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
