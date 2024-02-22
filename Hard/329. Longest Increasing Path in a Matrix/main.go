package main

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(matrix, visited, i, j))
		}
	}
	return ans
}

func dfs(matrix, visited [][]int, i, j int) int {
	if visited[i][j] > 0 {
		return visited[i][j]
	}
	if i > 0 && matrix[i][j] < matrix[i-1][j] {
		visited[i][j] = max(visited[i][j], dfs(matrix, visited, i-1, j))
	}
	if i < len(matrix)-1 && matrix[i][j] < matrix[i+1][j] {
		visited[i][j] = max(visited[i][j], dfs(matrix, visited, i+1, j))
	}
	if j > 0 && matrix[i][j] < matrix[i][j-1] {
		visited[i][j] = max(visited[i][j], dfs(matrix, visited, i, j-1))
	}
	if j < len(matrix[0])-1 && matrix[i][j] < matrix[i][j+1] {
		visited[i][j] = max(visited[i][j], dfs(matrix, visited, i, j+1))
	}
	visited[i][j]++
	return visited[i][j]
}
