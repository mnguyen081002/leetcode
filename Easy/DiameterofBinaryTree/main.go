package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxdiameter int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)

	maxdiameter = max(left, right)

	return 1 + maxdiameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)

	maxdiameter = max(maxdiameter, left+right)

	return 1 + max(left+right)
}
