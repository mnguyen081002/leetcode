package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	head, sum := 0, 0
	length := len(nums) + 1

	for tail := 0; tail < len(nums); tail++ {
		sum += nums[tail]
		for sum >= target {
			if tail-head+1 < length {
				length = tail - head + 1
			}
			sum -= nums[head]
			head++
		}
	}
	if length == len(nums)+1 {
		return 0
	}
	return length
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}

	fmt.Println(minSubArrayLen(7, nums))
}
