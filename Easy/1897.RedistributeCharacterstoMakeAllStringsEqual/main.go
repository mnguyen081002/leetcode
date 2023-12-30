package main

func makeEqual(words []string) bool {
	wordsCount := len(words)
	charCount := make(map[rune]int)

	for _, word := range words {
		for _, char := range word {
			charCount[char]++
		}
	}

	for _, value := range charCount {
		if value%wordsCount != 0 {
			return false
		}
	}

	return true
}
