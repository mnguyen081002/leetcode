package main

var directions = [][]int{{1, 0}, {0, -1}, {0, 1}, {-1, 0}}

func getMaximumGold(grid [][]int) int {
	ans, visit := 0, make([][]bool, len(grid))
	for i := range grid {
		visit[i] = make([]bool, len(grid[0]))
	}
	var backtracking func(int, int, int)
	backtracking = func(r, c, gold int) {
		if grid[r][c] == 0 {
			return
		}
		visit[r][c] = true
		gold += grid[r][c]
		ans = max(ans, gold)

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && !visit[nr][nc] {
				backtracking(nr, nc, gold)
			}
		}
		visit[r][c] = false
	}

	for i := range grid {
		for j := range grid[i] {
			backtracking(i, j, 0)
		}
	}

	return ans
}

func main() {
	grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	println(getMaximumGold(grid))
}
