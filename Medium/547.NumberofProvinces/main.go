package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	graph := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if isConnected[i-1][j-1] == 1 {
				graph[i] = append(graph[i], j)
			}
		}
	}

	visited := make([]bool, n+1)
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for _, v := range graph[i] {
			if visited[v] {
				continue
			}
			dfs(v)
		}
	}

	p := 0
	for i := 1; i <= n; i++ {
		if !visited[i] {
			p++
			dfs(i)
		}
	}
	return p
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}
