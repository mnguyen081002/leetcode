package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	maxLevel, ans := 0, 0
	dfs(root, 0, &maxLevel, &ans)
	return ans
}

func dfs(root *TreeNode, level int, maxLevel, ans *int) {
	if root == nil {
		return
	}
	if level == *maxLevel {
		*ans = root.Val
		*maxLevel++
	}
	dfs(root.Left, level+1, maxLevel, ans)
	dfs(root.Right, level+1, maxLevel, ans)
}
