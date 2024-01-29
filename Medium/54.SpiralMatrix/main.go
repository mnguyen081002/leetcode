package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	r, mr, c, mc := 0, len(matrix)-1, 0, len(matrix[0])-1
	rs := []int{}
	for {
		rs = append(rs, matrix[r][c:mc+1]...)
		r++
		if r > mr {
			break
		}
		for row := r; row <= mr; row++ {
			rs = append(rs, matrix[row][mc])
		}
		mc--
		if c > mc {
			break
		}
		for col := mc; col >= c; col-- {
			rs = append(rs, matrix[mr][col])
		}
		mr--
		if r > mr {
			break
		}
		for row := mr; row >= r; row-- {
			rs = append(rs, matrix[row][c])
		}
		c++
		if c > mc {
			break
		}
	}
	return rs
}

func main() {
	a := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	fmt.Println(spiralOrder(a))
}
