package main

func characterReplacement(s string, k int) int {
	charcounts := make([]int, 26)
	low, maxCounter, maxLen := 0, 0, 0

	for high := 0; high < len(s); high++ {
		charcounts[s[high]-'A']++
		maxCounter = max(maxCounter, charcounts[s[high-'A']])

		if maxCounter+k >= high-low+1 {
			maxLen = max(maxLen, high-low+1)
		} else {
			charcounts[s[low]-'A']--
			low++
		}
	}

	return maxLen
}

func main() {
	s := "AABABBA"
	characterReplacement(s, 1)
}
