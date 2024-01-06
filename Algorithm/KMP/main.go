package main

import "fmt"

func buildPrefixTable(pattern string) []int {
	m := len(pattern)
	table := make([]int, m)
	table[0] = 0
	j := 0

	for i := 1; i < m; {
		if pattern[i] == pattern[j] {
			j++
			table[i] = j
			i++
		} else {
			if j != 0 {
				j = table[j-1]
			} else {
				table[i] = 0
				i++
			}
		}
	}

	return table
}

func kmpSearch(text, pattern string) []int {
	n, m := len(text), len(pattern)
	table := buildPrefixTable(pattern)
	result := []int{}

	i, j := 0, 0
	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			result = append(result, i-j)
			j = table[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = table[j-1]
			} else {
				i++
			}
		}
	}

	return result
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"
	indices := kmpSearch(text, pattern)

	if len(indices) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", indices)
	} else {
		fmt.Println("Pattern not found in the text.")
	}
}
