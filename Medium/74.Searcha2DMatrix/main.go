package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n := len(matrix[0])
	m := len(matrix)
	l, r := 0, m*n-1

	for l <= r {
		mid := l + (r-l)/2
		row := mid / n
		col := mid - row*n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(11 / 3)
}
