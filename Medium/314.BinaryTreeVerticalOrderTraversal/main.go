package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeCol struct {
	Node *TreeNode
	Col  int
}

func verticalOrder(root *TreeNode) [][]int {
	minCol, maxCol := 0, 0
	var ccol func(*TreeNode, int)
	ccol = func(root *TreeNode, col int) {
		if root == nil {
			return
		}
		maxCol = max(maxCol, col)
		minCol = min(minCol, col)
		ccol(root.Left, col-1)
		ccol(root.Right, col+1)
	}
	ccol(root, 0)
	res := make([][]int, maxCol-minCol+1)

	q := []TreeCol{{root, 0}}
	for len(q) > 0 {
		f := q[0]
		q = q[1:]
		if f.Node == nil {
			continue
		}
		res[f.Col-minCol] = append(res[f.Col-minCol], f.Node.Val)
		if f.Node.Right != nil {
			q = append(q, TreeCol{Node: f.Node.Right, Col: f.Col + 1})
		}
		if f.Node.Left != nil {
			q = append(q, TreeCol{Node: f.Node.Left, Col: f.Col - 1})
		}
	}
	return res
}

func main() {
	verticalOrder(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}})
}
