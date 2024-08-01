package main

func equalPairs(grid [][]int) int {
	count := 0
	for r := range grid {
		for c := range grid {
			isD := false
			for i := range grid {
				if grid[r][i] != grid[i][c] {
					isD = true
					break
				}
			}
			if !isD {
				count++
			}
		}
	}

	return count
}

func main() {
	equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}})
}
