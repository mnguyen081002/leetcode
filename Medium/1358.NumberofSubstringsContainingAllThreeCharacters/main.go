package main

func numberOfSubstrings(s string) int {
	res, last := 0, [3]int{-1, -1, -1}
	for i := range s {
		last[s[i]-'a'] = i
		res += 1 + min(last[0], last[1], last[2])
	}
	return res
}

func main() {
	s := "abcabc"
	println(numberOfSubstrings(s))
}
