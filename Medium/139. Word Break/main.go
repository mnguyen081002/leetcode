package main

func wordBreak(s string, wordDict []string) bool {
	t := Constructor()
	for _, w := range wordDict {
		t.Insert(w)
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := range s {
		if !dp[i] {
			continue
		}
		for _, ln := range t.SearchSubstrings(s[i:]) {
			dp[i+ln] = true
		}
	}
	return dp[n]
}

type Trie struct {
	children   [28]*Trie
	isTerminal bool
}

func Constructor() Trie {
	return Trie{[28]*Trie{}, false}
}
func (this *Trie) Insert(word string) {
	for _, letter := range word {
		idx := letter - 'a'
		if this.children[idx] == nil {
			trie := Constructor()
			this.children[idx] = &trie
		}
		this = this.children[idx]
	}
	this.isTerminal = true
}
func (this *Trie) SearchSubstrings(line string) []int {
	result := []int{}
	for i, letter := range line {
		idx := letter - 'a'
		if this == nil {
			break
		}
		this = this.children[idx]

		if this != nil && this.isTerminal {
			result = append(result, i+1)
		}
	}
	return result
}
