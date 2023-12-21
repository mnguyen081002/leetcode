package main

import "fmt"

func longestPalindrome(s string) int {
	count := 0
	m := map[rune]uint8{}
	for _, v := range s {
		if _, found := m[v]; found {
			delete(m, v)
			count += 2
			continue
		}
		m[v] = 0
		fmt.Println(m)
	}
	if len(m) != 0 {
		return count + 1
	}
	return count
}

func main() {
	longestPalindrome("abccccdd")
}
