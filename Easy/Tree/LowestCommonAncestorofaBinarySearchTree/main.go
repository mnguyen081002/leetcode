package main

import (
	"fmt"
	common "leetcode/Easy/Tree/Common"
)

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func main() {
	values := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	var root *common.TreeNode
	for _, value := range values {
		root = common.Insert(root, value)
	}

	r := lowestCommonAncestor(root, &common.TreeNode{Val: 9}, &common.TreeNode{Val: 4})

	common.PrintPretty(root, "", true)

	fmt.Println(r.Val)
}
