package main

import (
	"fmt"
	"slices"
)

func customSortString(order string, s string) string {
	m := map[rune]int{}
	for i, v := range order {
		m[v] = 26 - i
	}
	x := []rune(s)
	slices.SortFunc(x, func(x, y rune) int {
		return m[y] - m[x]
	})

	return string(x)
}

func main() {
	fmt.Println(customSortString("cba", "abcd"))
}
