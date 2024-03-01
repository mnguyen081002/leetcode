package main

func findPeakElement(nums []int) int {
	var finder func(int, int) int
	finder = func(left int, right int) int {
		if left == right {
			return left
		}
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			return finder(left, mid)
		} else {
			return finder(mid+1, right)
		}
	}
	return finder(0, len(nums)-1)
}
