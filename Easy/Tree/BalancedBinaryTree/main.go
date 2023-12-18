package main

import (
	"fmt"
	common "leetcode/Easy/Tree/Common"
)

func isBalanced(root *common.TreeNode) bool {
	s, _ := dfs(root)
	return s
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func dfs(root *common.TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBanlanced, rightHeight := dfs(root.Right)

	diff := abs(leftHeight - rightHeight)

	if isLeftBalanced && isRightBanlanced && diff < 2 {
		return true, 1 + max(leftHeight, rightHeight)
	}

	return false, 0
}

func main() {
	root := &common.TreeNode{
		Val: 1,
		Right: &common.TreeNode{
			Val: 2,
			Right: &common.TreeNode{
				Val: 3,
				Right: &common.TreeNode{
					Val: 4,
					Right: &common.TreeNode{
						Val: 5,
					},
				},
			},
			Left: &common.TreeNode{
				Val: 3,
			},
		},
		Left: &common.TreeNode{
			Val: 2,
		},
	}
	r := isBalanced(root)

	common.PrintPretty(root, "", true)

	fmt.Println(r)
}
