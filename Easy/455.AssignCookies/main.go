package main

import "sort"

func findContentChildren(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)

	for i := 0; i < len(s) && res < len(g); i++ {
		if s[i] >= g[res] {
			res++
		}
	}

	return res
}
