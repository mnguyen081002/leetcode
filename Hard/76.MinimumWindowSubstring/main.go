package main

import "math"

func minWindow(s string, t string) string {
	charmap := make([]int, 128)
	for i := 0; i < len(t); i++ {
		charmap[t[i]]++
	}
	end, begin, head, d, counter := 0, 0, 0, math.MaxInt, len(t)
	for end < len(s) {
		if charmap[s[end]] > 0 {
			counter--
		}
		charmap[s[end]]--
		end++

		for counter == 0 {
			if end-begin < d {
				head = begin
				d = end - head
			}
			charmap[s[begin]]++
			if charmap[s[begin]] > 0 {
				counter++
			}
			begin++
		}
	}

	if d == math.MaxInt {
		return ""
	}

	return s[head : head+d]
}

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}
