package main

func dfs(u int, visited []bool, adj [][]int) {
	visited[u] = true
	for _, v := range adj[u] {
		if !visited[v] {
			dfs(v, visited, adj)
		}
	}
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	adj := make([][]int, n)
	for _, edge := range connections {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	visited := make([]bool, n)
	var count int = 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i, visited, adj)
		}
	}
	return count - 1
}
