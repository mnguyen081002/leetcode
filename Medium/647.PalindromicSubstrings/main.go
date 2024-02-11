package main

func countSubstrings(s string) int {
	var count = 0
	for i := 0; i < len(s); i++ {
		var l, r = i, i
		count++ // single letter
		for 0 <= l-1 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
			count++
		}
		l, r = i+1, i
		for 0 <= l-1 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
			count++
		}
	}
	return count
}

func main() {
	countSubstrings("aaa")
}
