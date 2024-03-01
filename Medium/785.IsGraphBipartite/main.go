package main

func isBipartite(graph [][]int) bool {
	visited := make([]int, len(graph))
	for i := range visited {
		if visited[i] != 0 {
			continue
		}
		if !dfs(i, 1, graph, visited) {
			return false
		}
	}
	return true
}
func dfs(i, v int, graph [][]int, visited []int) bool {
	for _, vv := range graph[i] {
		if visited[vv] == 0 {
			visited[vv] = v
			if !dfs(vv, -v, graph, visited) {
				return false
			}
			continue
		}
		if visited[vv] != v {
			return false
		}
	}
	return true
}
