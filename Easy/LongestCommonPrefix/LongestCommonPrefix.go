package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	prefix := ""
	length := len(strs)
	if length == 0 {
		return ""
	}
	for ci, c := range strs[0] {
		for i := 1; i < length; i++ {
			st := strs[i]
			if ci > len(st)-1 {
				return prefix
			}
			if c != rune(st[ci]) {
				return prefix
			}
		}
		prefix += string(c)
	}

	return prefix
}

func main() {
	strs := []string{"helli", "helli", "heli"}
	fmt.Println(longestCommonPrefix(strs))
}
