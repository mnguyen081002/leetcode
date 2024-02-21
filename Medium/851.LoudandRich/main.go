package main

func loudAndRich(richer [][]int, quiet []int) []int {
	graph := make([][]int, len(quiet))
	for _, e := range richer {
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	ans := make([]int, len(quiet))
	for i := range quiet {
		ans[i] = -1
	}

	var dfs func(int) int

	dfs = func(node int) int {
		if ans[node] != -1 {
			return ans[node]
		}
		mi := node
		for _, n := range graph[node] {
			m := dfs(n)
			if quiet[mi] > quiet[m] {
				mi = m
			}
		}
		ans[node] = mi
		return mi
	}
	for n := range quiet {
		dfs(n)
	}
	return ans
}
