package main

func findRedundantConnection(edges [][]int) []int {
	roots := make([]int, len(edges)+1)

	for i := range roots {
		roots[i] = i
	}

	for _, e := range edges {
		if find(e[0], roots) == find(e[1], roots) {
			return []int{e[0], e[1]}
		}
		union(e[0], e[1], roots)
	}

	return nil
}

func find(x int, roots []int) int {
	if roots[x] == x {
		return x
	}
	root := find(roots[x], roots)
	roots[x] = root
	return root
}

func union(x, y int, roots []int) {
	roots[find(y, roots)] = find(x, roots)
}
