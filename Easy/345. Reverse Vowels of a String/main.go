package main

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	result := []byte(s)
	for l < r {
		for !isVowels(s[l]) && l < r {
			l++
		}

		for !isVowels(s[r]) && l < r {
			r--
		}

		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}

func isVowels(a byte) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u'
}
