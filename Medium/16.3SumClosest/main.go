package main

import (
	"fmt"
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	closest := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if min(abs(sum-target), abs(closest-target)) != abs(closest-target) {
				closest = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	if closest == math.MaxInt {
		return 0
	}

	return closest
}

func abs(n1 int) int {
	if n1 < 0 {
		return -n1
	}
	return n1
}

func main() {
	nums := []int{2, 3, 8, 9, 10}
	fmt.Println(threeSumClosest(nums, 16))
}
