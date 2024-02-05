package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	r, l := dfs(root.Right, p, q), dfs(root.Left, p, q)
	if r != nil && l == nil {
		return r
	} else if r == nil && l != nil {
		return l
	} else if r == nil && l == nil {
		return nil
	} else {
		return root
	}
}
