package main

import (
	"fmt"
)

// 1. s = "1", ss, counter, pre := "", '1', rune(s[0])
// do nothing
// 2. s =
func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		ss, counter, pre := "", '0', '-'
		for i := 0; i < len(s); i++ {
			if pre == '-' {
				pre = rune(s[i])
				counter = '1'
				continue
			}

			if rune(s[i]) != pre {
				ss += string(counter) + string(pre)
				pre = rune(s[i])
				counter = '1'
			} else {
				counter++
			}
		}

		ss += string(counter) + string(pre)
		s = ss
	}

	return s
}

// countAndSay(1) = "1"
// countAndSay(2) = say "1" = one 1 = "11"
// countAndSay(3) = say "11" = two 1's = "21"
// countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
// countAndSay(5) = say "1211" = one 1 + one two + two one = "111221"

func main() {
	fmt.Println(countAndSay(4))
}
