package main

import (
	"container/heap"
	"math"
)

func countRestrictedPaths(n int, edges [][]int) int {
	graph := make([][][]int, n+1)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], []int{w, v})
		graph[v] = append(graph[v], []int{w, u})
	}

	dist := dijkstra(n, graph)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return dfs(1, n, graph, dist, memo)
}

func dfs(src, dst int, graph [][][]int, dist, memo []int) int {
	if memo[src] != -1 {
		return memo[src]
	}
	if src == dst {
		return 1
	}
	ans := 0
	for _, nei := range graph[src] {
		v := nei[1]
		if dist[src] > dist[v] {
			ans = (ans + dfs(v, dst, graph, dist, memo)) % 1000000007
		}
	}

	memo[src] = ans

	return memo[src]
}

func dijkstra(n int, graph [][][]int) []int {
	q := &PriorityQueue{[2]int{0, n}}
	dist := make([]int, n+1)

	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[n] = 0

	for len(*q) > 0 {
		top := heap.Pop(q).([2]int)
		w, u := top[0], top[1]

		if w != dist[u] {
			continue
		}
		for _, e := range graph[u] {
			w, v := e[0], e[1]
			if dist[v] > w+dist[u] {
				dist[v] = w + dist[u]
				heap.Push(q, [2]int{dist[v], v})
			}
		}
	}

	return dist
}

type PriorityQueue [][2]int

func (q PriorityQueue) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func (q *PriorityQueue) Push(v interface{}) {
	*q = append(*q, v.([2]int))
}

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i][0] < q[j][0]
}
