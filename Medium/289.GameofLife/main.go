package main

const (
	Die     = 0
	Live    = 1
	LiveDie = 2
	DieLive = 3
)

var directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func gameOfLife(board [][]int) {

	for r := range board {
		for c := range board[0] {
			handleBoard(r, c, board)
		}
	}

	for r := range board {
		for c := range board[0] {
			if board[r][c] == LiveDie {
				board[r][c] = Die
			}
			if board[r][c] == DieLive {
				board[r][c] = Live
			}
		}
	}
}

func handleBoard(r, c int, board [][]int) {
	n, m := len(board), len(board[0])

	live := 0
	for _, d := range directions {
		rr, cc := r+d[0], c+d[1]

		if rr >= 0 && rr < n && cc >= 0 && cc < m {
			cell := board[rr][cc]
			if cell == Live || cell == LiveDie {
				live++
			}
		}
	}

	curr := board[r][c]
	if live < 2 || live > 3 {
		if curr == Live {
			board[r][c] = LiveDie
		}
	} else if live == 3 {
		if curr == Die {
			board[r][c] = DieLive
		}
	}
}

func main() {
	gameOfLife([][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	})
}
