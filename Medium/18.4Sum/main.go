package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, start, target int) [][]int {
	sum := 0
	result := [][]int{}
	end := len(nums) - 1
	for start < end {
		sum = nums[start] + nums[end]
		if sum == target {
			result = append(result, []int{nums[start], nums[end]})
			start++
			end--
		} else if sum < target {
			start++
		} else {
			end--
		}
	}

	return result
}

func kSum(nums []int, k, start, target int) [][]int {
	result := [][]int{}
	if start > len(nums)-1 || k < 2 || nums[start]*k > target || nums[len(nums)-1]*k < target {
		return [][]int{}
	}

	if k == 2 {
		return twoSum(nums, start, target)
	}

	for i := start; i < len(nums); i++ {
		for _, res := range kSum(nums, k-1, i+1, target-nums[i]) {
			result = append(result, append([]int{nums[i]}, res...))
		}
	}

	return result
}

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	return kSum(nums, 4, 0, target)
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}
