package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	sum2 := 0
	for i, v := range nums {
		sum2 += v
		if sum == sum2 {
			return i
		}
		sum -= v
	}

	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
