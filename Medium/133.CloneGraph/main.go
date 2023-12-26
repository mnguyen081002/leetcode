package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var visited = map[*Node]*Node{}

	var dfs func(*Node) *Node

	dfs = func(node *Node) *Node {
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		clonedNode := &Node{node.Val, []*Node{}}

		visited[node] = clonedNode

		if len(node.Neighbors) > 0 {
			for _, n := range node.Neighbors {
				clonedNode.Neighbors = append(clonedNode.Neighbors, dfs(n))
			}
		}

		return clonedNode
	}

	return dfs(node)
}
