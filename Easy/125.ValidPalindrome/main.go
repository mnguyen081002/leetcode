package main

import "unicode"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	left, right := 0, len(s)-1

	for left <= right {
		if f(rune(s[left])) == -1 {
			left++
			continue
		}
		if f(rune(s[right])) == -1 {
			right--
			continue
		}
		if f(rune(s[left])) != f(rune(s[right])) {
			return false
		}
		right--
		left++
	}

	return true
}
