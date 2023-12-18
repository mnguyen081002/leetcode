package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func printPretty(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	if root.Right != nil {
		printPretty(root.Right, prefix+func(isLeft bool) string {
			if isLeft {
				return "│   "
			}
			return "    "
		}(isLeft), false)
	}

	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Print("└───")
	} else {
		fmt.Print("┌───")
	}

	fmt.Println(root.Val)

	if root.Left != nil {
		printPretty(root.Left, prefix+func(isLeft bool) string {
			if isLeft {
				return "    "
			}
			return "│   "
		}(isLeft), true)
	}
}
func insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value, Left: nil, Right: nil}
	}

	if value < root.Val {
		root.Left = insert(root.Left, value)
	} else if value > root.Val {
		root.Right = insert(root.Right, value)
	}

	return root
}

func main() {
	values := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	var root *TreeNode
	for _, value := range values {
		root = insert(root, value)
	}

	r := lowestCommonAncestor(root, &TreeNode{Val: 9}, &TreeNode{Val: 4})

	printPretty(root, "", true)

	fmt.Println(r.Val)
}
