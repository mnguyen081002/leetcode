package main

import (
	"fmt"
	"strconv"
)

// TODO: handle this
func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}

// func findNthDigit(n int) int {
//     if n < 10 {
//         return n
//     }

//     if n % 2 == 0 {
//         return (n - 10) / 20 + 1
//     }else {
//         if n / 10 % 2 == 0 {
//             return (n % 10 + 10) / 2
//         }
//         return n % 10 / 2
//     }
// }

func main() {
	fmt.Println(1 / 10)
}
