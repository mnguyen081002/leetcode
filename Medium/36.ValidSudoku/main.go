package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidRow(board[i]) {
			return false
		}

		if !isValidCol(board, i) {
			return false
		}

		if !isValidCol(board, i) || !isValidSubBox(board, i/3*3, i%3*3) {
			return false
		}
	}

	return true
}

func isValidRow(row []byte) bool {
	c := make(map[byte]int)

	for i := 0; i < 9; i++ {
		if row[i] == '.' {
			continue
		}

		c[row[i]]++
	}

	for _, v := range c {
		if v > 1 {
			return false
		}
	}

	return true
}

func isValidCol(board [][]byte, col int) bool {
	c := make(map[byte]int)

	for i := 0; i < 9; i++ {
		if board[i][col] == '.' {
			continue
		}

		c[board[i][col]]++
	}

	for _, v := range c {
		if v > 1 {
			return false
		}
	}

	return true
}

func isValidSubBox(board [][]byte, row int, col int) bool {
	c := make(map[byte]int)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][col+j] == '.' {
				continue
			}

			c[board[row+i][col+j]]++
		}
	}

	for _, v := range c {
		if v > 1 {
			return false
		}
	}

	return true
}
