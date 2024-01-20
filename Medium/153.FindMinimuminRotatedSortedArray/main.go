package main

import "math"

func findMin(nums []int) int {
	m := math.MaxInt
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		m = min(m, nums[mid])

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return m
}
