package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}

	var queue []*TreeNode
	queue = append(queue, root)
	currentDepth := 1
	for len(queue) > 0 {
		currentDepth++
		length := len(queue)
		for i := 0; i < length; i++ {
			if currentDepth == depth {
				queue[i].Left = &TreeNode{val, queue[i].Left, nil}
				queue[i].Right = &TreeNode{val, nil, queue[i].Right}
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return root
}
