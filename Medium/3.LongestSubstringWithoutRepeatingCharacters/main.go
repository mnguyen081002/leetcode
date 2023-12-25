package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}
	var result, start int

	for i, v := range s {
		if mv, found := m[v]; found && mv >= start {
			start = mv + 1
		}
		m[v] = i
		result = max(i-start+1, result)
	}

	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring(" "))
}
