package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	result := 0
	for len(queue) > 0 {
		valSum := 0
		for _, _ = range queue {
			node := queue[0]
			valSum += node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = valSum
	}
	return result
}

func main() {
	deepestLeavesSum(&TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, &TreeNode{7, nil, nil}, nil}, &TreeNode{5, nil, nil},
		}, &TreeNode{3, nil,
			&TreeNode{6, nil, &TreeNode{8, nil, nil}},
		},
	})
}
