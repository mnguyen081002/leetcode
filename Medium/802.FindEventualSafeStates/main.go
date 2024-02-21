package main

import "slices"

func eventualSafeNodes(graph [][]int) []int {
	visit := map[int]int{}
	ans := []int{}

	var dfs func(int) bool
	dfs = func(node int) bool {
		if v, ok := visit[node]; ok {
			return v == 1
		}

		visit[node] = 0

		for _, n := range graph[node] {
			if !dfs(n) {
				return false
			}
		}

		visit[node] = 1
		ans = append(ans, node)
		return true
	}

	for i := 0; i < len(graph); i++ {
		dfs(i)
	}
	slices.Sort(ans)
	return ans
}
