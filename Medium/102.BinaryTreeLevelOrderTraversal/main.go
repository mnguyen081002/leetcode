package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = [][]int{}
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(result)-1 < level {
		result = append(result, []int{root.Val})
	} else {
		result[level] = append(result[level], root.Val)
	}

	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
