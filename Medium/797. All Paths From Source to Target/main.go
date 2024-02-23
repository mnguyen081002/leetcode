package main

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	var dfs func(int, []int)
	dfs = func(node int, curr []int) {
		if node == len(graph)-1 {
			ans = append(ans, append([]int{}, curr...))
			return
		}

		for _, e := range graph[node] {
			dfs(e, append(curr, e))
		}
	}
	dfs(0, []int{0})
	return ans
}
