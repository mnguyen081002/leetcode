package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	dfs(root, m)
	return find(m)
}

func dfs(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}

	sum := root.Val
	sum += dfs(root.Left, m) + dfs(root.Right, m)
	m[sum]++

	return sum
}

func find(m map[int]int) []int {
	r := []int{}
	max := 0
	for key, count := range m {
		if count > max {
			max = count
			r = nil
			r = append(r, key)
		} else if count == max {
			r = append(r, key)
		}
	}
	return r
}
