package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var y int

	for x > y {
		y = x%10 + y*10
		if x == y {
			return true
		}
		x /= 10
	}
	return x == y
}

func main() {
	fmt.Println(isPalindrome(121))
}
