package main

func solveSudoku(board [][]byte) {
	rm := make([]int, 9)
	cm := make([]int, 9)
	grid := make([]int, 9)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				set(board, board[r][c]-'0', r, c, rm, cm, grid)
			}
		}
	}

	backtracking(board, 0, 0, rm, cm, grid)
}

// 872 in binary is 1101101000
func canSet(value byte, r, c int, rm, cm []int, grid []int) bool {
	gi := r/3*3 + c/3

	if isSet(value, gi, grid) {
		return false
	}
	if isSet(value, r, rm) {
		return false
	}
	if isSet(value, c, cm) {
		return false
	}

	return true
}

// 170 in binary is 10101010 so that means  2

func set(board [][]byte, value byte, r, c int, rm, cm []int, grid []int) {
	rm[r] |= (1 << (value - 1))
	cm[c] |= (1 << (value - 1))
	board[r][c] = value + '0'

	gi := r/3*3 + c/3
	grid[gi] |= (1 << (value - 1)) // ex: 1 << 9 = 1000000000 then 1 << 2 = 1000000000 | 100 = 1000000100
}

func unset(board [][]byte, value byte, r, c int, rm, cm []int, grid []int) {
	rm[r] ^= (1 << (value - 1))
	cm[c] ^= (1 << (value - 1))
	gi := r/3*3 + c/3
	grid[gi] ^= (1 << (value - 1))
	board[r][c] = '.'
}

func backtracking(board [][]byte, r, c int, rm, cm []int, grid []int) bool {
	if r == 9 {
		return true
	}

	if c == 9 {
		return backtracking(board, r+1, 0, rm, cm, grid)
	}

	if board[r][c] != '.' {
		return backtracking(board, r, c+1, rm, cm, grid)
	}

	for i := 1; i <= 9; i++ {
		if canSet(byte(i), r, c, rm, cm, grid) {
			set(board, byte(i), r, c, rm, cm, grid)
			if backtracking(board, r, c+1, rm, cm, grid) {
				return true
			}
			unset(board, byte(i), r, c, rm, cm, grid)
		}
	}

	return false
}

func isSet(value byte, x int, m []int) bool {
	return (m[x] & (1 << (value - 1))) != 0
}
