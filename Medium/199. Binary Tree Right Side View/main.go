package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res, 0)
	return res
}

func dfs(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if len(*res)-1 < level {
		*res = append(*res, root.Val)
	} else {
		(*res)[level] = root.Val
	}
	dfs(root.Left, res, level+1)
	dfs(root.Right, res, level+1)
}
