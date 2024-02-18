package main

import "fmt"

var directions [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func numIslands(grid [][]byte) int {
	count := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				count++
				traverse(grid, row, col)
			}
		}
	}

	return count
}

func traverse(grid [][]byte, row, col int) {
	if grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	rows, cols := len(grid), len(grid[0])
	for _, direction := range directions {
		next_r, next_c := row+direction[0], col+direction[1]
		if 0 <= next_r && next_r < rows && 0 <= next_c && next_c < cols {
			traverse(grid, next_r, next_c)
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Println("Number of islands:", numIslands(grid))
}
