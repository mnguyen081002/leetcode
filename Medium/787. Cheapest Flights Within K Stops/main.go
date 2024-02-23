package main

import "math"

type Flights struct {
	dst   int
	price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]Flights, n)

	for _, v := range flights {
		graph[v[0]] = append(graph[v[0]], Flights{v[1], v[2]})
	}

	var dfs func(int, int, int, int) int
	dfs = func(src, dst, k, cost int) int {
		if src == dst {
			return cost
		}
		if k == 0 {
			return -1
		}
		m := math.MaxInt

		for _, e := range graph[src] {
			m = min(m, dfs(e.dst, dst, k-1, cost+e.price))
		}
		return m
	}
	return dfs(src, dst, k, 0)
}
