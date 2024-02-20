package main

func pacificAtlantic(heights [][]int) [][]int {
	res := [][]int{}
	ROWS, COLS := len(heights), len(heights[0])
	pacific, atlantic := make([][]bool, ROWS), make([][]bool, ROWS)

	for r := 0; r < ROWS; r++ {
		pacific[r] = make([]bool, COLS)
		atlantic[r] = make([]bool, COLS)
	}

	for r := 0; r < ROWS; r++ {
		dfs(heights, heights[r][0], pacific, r, 0)
		dfs(heights, heights[r][COLS-1], atlantic, r, COLS-1)
	}

	for c := 0; c < COLS; c++ {
		dfs(heights, heights[0][c], pacific, 0, c)
		dfs(heights, heights[ROWS-1][c], atlantic, ROWS-1, c)
	}

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

func dfs(heights [][]int, height int, visit [][]bool, r, c int) {
	if r == len(heights) || c == len(heights[0]) || c < 0 || r < 0 || heights[r][c] < height || visit[r][c] {
		return
	}
	visit[r][c] = true
	dfs(heights, heights[r][c], visit, r, c+1)
	dfs(heights, heights[r][c], visit, r, c-1)
	dfs(heights, heights[r][c], visit, r+1, c)
	dfs(heights, heights[r][c], visit, r-1, c)
}
