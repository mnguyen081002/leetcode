package main

import "math"

func searchRange(nums []int, target int) []int {
	mn, mx := math.MaxInt, -1
	binarySearch(nums, 0, len(nums)-1, target, &mn, &mx)
	if mn == math.MaxInt {
		if mx != -1 {
			mn = mx
		} else {
			mn = -1
		}
	}
	if mx == -1 {
		if mn != math.MaxInt {
			mx = mn
		}
	}
	return []int{mn, mx}
}

func binarySearch(nums []int, left, right, target int, mn, mx *int) {
	if left > right || len(nums) == 0 {
		return
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		*mn = min(*mn, mid)
		*mx = max(*mx, mid)
		binarySearch(nums, left, mid-1, target, mn, mx)
		binarySearch(nums, mid+1, right, target, mn, mx)
		return
	}

	if nums[mid] > target {
		binarySearch(nums, left, mid-1, target, mn, mx)
		return
	}
	binarySearch(nums, mid+1, right, target, mn, mx)
}
