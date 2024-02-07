package main

import "math"

// testcase: '[5,4,6,1,10,3,7,null,null,null,null,1,2]'

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val < max && root.Val > min {
		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}
	return false
}
