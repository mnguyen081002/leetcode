package main

func findJudge(n int, trust [][]int) int {
	graph := make([]int, n)
	htrust := map[int]int{}
	for _, e := range trust {
		graph[e[1]-1]++
		htrust[e[0]-1] = 1
	}

	for i, e := range graph {
		if e == n-1 {
			if _, f := htrust[i]; !f {
				return i + 1
			}
		}
	}

	return -1
}
