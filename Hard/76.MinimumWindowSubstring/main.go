package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	charMap := make([]int, 128)
	for _, c := range t {
		charMap[c]++
	}

	counter, begin, end, d, head := len(t), 0, 0, math.MaxInt32, 0

	for end < len(s) {
		if charMap[s[end]] > 0 {
			counter-- // in t
		}
		charMap[s[end]]--
		end++

		for counter == 0 { // valid
			if end-begin < d {
				head = begin
				d = end - head
			}
			charMap[s[begin]]++
			if charMap[s[begin]] > 0 {
				counter++ // make it invalid
			}
			begin++
		}
	}

	if d == math.MaxInt32 {
		return ""
	}
	return s[head : head+d]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
