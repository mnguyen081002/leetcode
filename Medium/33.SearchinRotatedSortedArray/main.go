package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[high] {
			if nums[mid] > target && nums[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
