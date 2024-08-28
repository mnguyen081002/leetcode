package main

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	count := 0

	for r := range grid2 {
		for c := range grid2[0] {
			if grid2[r][c] == 1 && grid1[r][c] == 1 {
				count++
				dfs(grid1, grid2, r, c)
			}
		}
	}

	return count
}

func dfs(grid1, grid2 [][]int, r, c int) {
	grid2[r][c] = 0
	grid1[r][c] = 0
	for _, d := range directions {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < len(grid2) && nc >= 0 && nc < len(grid2[0]) && grid2[nr][nc] == 1 {
			dfs(grid1, grid2, nr, nc)
		}
	}
}

func main() {
	grid1 := [][]int{{1, 1, 1, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	countSubIslands(grid1, grid2)
}
