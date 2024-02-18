package main

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func orangesRotting(grid [][]int) int {
	time := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 2 {
				rotting(grid, row, col, 0, &time)
			}
		}
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				return -1
			}
		}
	}

	return time
}

func rotting(grid [][]int, row, col, currentTime int, time *int) {
	if grid[row][col] == -1 || grid[row][col] == 0 {
		return
	}
	if grid[row][col] == 2 && currentTime > *time {
		currentTime = 0
		*time = *time / 2
	}
	grid[row][col] = -1
	*time = max(*time, currentTime)

	rows, cols := len(grid), len(grid[0])
	for _, direction := range directions {
		next_r, next_c := row+direction[0], col+direction[1]
		if 0 <= next_r && next_r < rows && 0 <= next_c && next_c < cols {
			rotting(grid, next_r, next_c, currentTime+1, time)
		}
	}
}

func main() {
	grid := [][]int{
		{2, 1, 0, 2},
	}

	println(orangesRotting(grid))
}
