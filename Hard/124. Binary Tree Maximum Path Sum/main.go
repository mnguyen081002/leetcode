package main

import "math"

// testcase: [-9,9,20,11,16,15,7]
// expected: 51
// [-9,9,20,11,16,15,7,null,null,10,5,null,null,null,null]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	mx := root.Val
	dfs(root, &mx)
	return mx
}

func dfs(root *TreeNode, mx *int) int {
	if root == nil {
		return math.MinInt
	}
	maxSL := max(dfs(root.Left, mx), 0)
	maxSR := max(dfs(root.Right, mx), 0)
	*mx = max(*mx, root.Val+maxSL+maxSR)
	return root.Val + max(maxSL, maxSR)
}
