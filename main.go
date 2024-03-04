package main

import "fmt"

// Count Submatrices with Top-Left Element and Sum Less Than k
func countSubmatrices(grid [][]int, k int) int {
	count := 0
	n, m, maxSize := len(grid), len(grid[0]), min(len(grid), len(grid[0]))
	for size := 1; size <= maxSize; size++ {
		for i := 0; i <= n-size; i++ {
			for j := 0; j <= m-size; j++ {
				sum := 0
				for x := 0; x < i+size; x++ {
					for y := 0; y < j+size; y++ {
						sum += grid[x][y]
					}
				}
				if sum <= k {
					count++
				}
			}
		}
	}
	return count
}

func main() {

	fmt.Println(countSubmatrices([][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20))
}
