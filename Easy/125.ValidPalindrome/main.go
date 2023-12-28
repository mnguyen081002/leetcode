package main

import "unicode"

func isNotValid(r byte) bool {
	return !unicode.IsLetter(rune(r)) && !unicode.IsNumber(rune(r))
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	left, right := 0, len(s)-1

	for left <= right {
		if isNotValid(s[left]) {
			left++
			continue
		}
		if isNotValid(s[right]) {
			right--
			continue
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		right--
		left++
	}

	return true
}
