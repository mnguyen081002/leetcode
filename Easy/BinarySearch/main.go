package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if right >= left {
		mid := (right + left) / 2
		v := nums[mid]
		if target == v {
			return mid
		}
		if target > v {
			return binarySearch(nums, mid+1, right, target)
		}
		return binarySearch(nums, left, mid-1, target)
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 2))
}
