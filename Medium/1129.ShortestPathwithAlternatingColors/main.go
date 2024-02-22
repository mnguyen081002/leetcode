package main

const (
	RED = iota
	BLUE
)

type state struct {
	node  int
	color int
	steps int
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	graph := make([][2][]int, n)
	for _, r := range redEdges {
		graph[r[0]][RED] = append(graph[r[0]][RED], r[1])
	}

	for _, b := range blueEdges {
		graph[b[0]][BLUE] = append(graph[b[0]][BLUE], b[1])
	}

	queue := []state{}
	queue = append(queue, state{node: 0, color: RED, steps: 0})
	queue = append(queue, state{node: 0, color: BLUE, steps: 0})

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	seen := make([][2]bool, n)
	seen[0][RED] = true
	seen[0][BLUE] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if ans[node.node] == -1 {
			ans[node.node] = node.steps
		}
		for _, next := range graph[node.node][node.color] {
			if !seen[next][1-node.color] {
				queue = append(queue, state{node: next, color: 1 - node.color, steps: node.steps + 1})
				seen[next][1-node.color] = true
			}
		}
	}
	return ans
}
