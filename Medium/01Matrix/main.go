package main

func updateMatrix(mat [][]int) [][]int {
	var direction = []int{0, 1, 0, -1, 0}
	m := len(mat)
	n := len(mat[0])

	queue := [][]int{}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if mat[r][c] == 0 {
				queue = append(queue, []int{r, c})
			} else {
				mat[r][c] = -1
			}
		}
	}

	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		r := p[0]
		c := p[1]
		for i := 0; i < 4; i++ {
			nr := r + direction[i]
			nc := c + direction[i+1]
			if nr == m || nr < 0 || nc == n || nc < 0 || mat[nr][nc] != -1 {
				continue
			}

			mat[nr][nc] = mat[r][c] + 1
			queue = append(queue, []int{nr, nc})
		}
	}

	return mat
}
