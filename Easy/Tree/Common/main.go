package common

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintPretty(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	if root.Right != nil {
		PrintPretty(root.Right, prefix+func(isLeft bool) string {
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
		PrintPretty(root.Left, prefix+func(isLeft bool) string {
			if isLeft {
				return "    "
			}
			return "│   "
		}(isLeft), true)
	}
}
func Insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value, Left: nil, Right: nil}
	}

	if value < root.Val {
		root.Left = Insert(root.Left, value)
	} else if value > root.Val {
		root.Right = Insert(root.Right, value)
	}

	return root
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Find the middle element of the array
	mid := len(nums) / 2

	// Create a new node with the middle element as the value
	node := &TreeNode{Val: nums[mid]}

	// Recursively build the left and right subtrees
	node.Left = SortedArrayToBST(nums[:mid])
	node.Right = SortedArrayToBST(nums[mid+1:])

	return node
}
