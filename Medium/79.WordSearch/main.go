package main

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	visit := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visit[i] = make([]bool, len(board[0]))
	}

	var backtrack func(int, int, string) bool
	backtrack = func(r, c int, word string) bool {
		if word == "" {
			return true
		}

		if board[r][c] != word[0] || visit[r][c] {
			return false
		}
		visit[r][c] = true
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < len(board) && nc < len(board[0]) && nr >= 0 && nc >= 0 && backtrack(nr, nc, word[1:]) {
				return true
			}
		}
		visit[r][c] = false
		return false
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if backtrack(r, c, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}

	exist(board, "ABCCED")
}
