package main

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	Fresh  = 1
	Rotten = 2
)

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var rotten [][2]int
	fresh := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == Fresh {
				fresh++
			}
			if grid[r][c] == Rotten {
				rotten = append(rotten, [2]int{r, c})
			}
		}
	}

	if fresh == 0 {
		return 0
	}
	minutes := 0
	for len(rotten) > 0 {
		for _, v := range rotten {
			for _, direction := range directions {
				nr, nc := v[0]+direction[0], v[1]+direction[1]
				if nr >= 0 && nc >= 0 && nr < m && nc < n && grid[nr][nc] == Fresh {
					fresh--
					grid[nr][nc] = Rotten
					rotten = append(rotten, [2]int{nr, nc})
				}
			}
			rotten = rotten[1:]
		}
		minutes++
		if fresh == 0 {
			return minutes
		}
	}
	return -1
}
