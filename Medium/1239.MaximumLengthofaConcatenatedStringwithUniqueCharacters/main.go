package main

import (
	"fmt"
)

func maxLength(arr []string) int {
	m := backtracking(arr, 0)
	return m
}

func backtracking(arr []string, bcurr int32) int {
	mx := 0
	for i, v := range arr {
		if y, bb := isValid(bcurr, v); y && len(arr[i:]) != 0 {
			mx = max(backtracking(arr[i:], bb), mx)
		} else {
			mx = max(mx, countSetBits(bcurr))
		}
	}

	return mx
}

func countSetBits(num int32) int {
	count := 0

	for num > 0 {
		// Kiểm tra bit cuối cùng có phải là 1 hay không
		if num&1 == 1 {
			count++
		}

		// Dịch phải số để kiểm tra bit tiếp theo
		num >>= 1
	}

	return count
}

// func isValid(bs1 int32, s2 string) bool {
// 	for _, v := range s2 {
// 		fmt.Print("bs1: ", strconv.FormatInt(int64(bs1), 2), " ")
// 		fmt.Print("s2: ", strconv.FormatInt(int64(v), 2))
// 		fmt.Print("\n")
// 		if bs1&(1<<(v-96)) != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

func isValid(b int32, s string) (bool, int32) {
	for _, v := range s {
		if b != 0 && b&(1<<(v-96)) != 0 {
			return false, 0
		}
		b |= (1 << (v - 96))
	}
	return true, b
}

func main() {
	arr := []string{"cha", "r", "act", "ers"}
	fmt.Println(maxLength(arr))
}
