package main

import "slices"

func divideArray(nums []int, k int) [][]int {
	ans := make([][]int, len(nums)/3)
	slices.Sort(nums)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		ans[i/3] = []int{nums[i], nums[i+1], nums[i+2]}
	}
	return ans
}

func main() {
	nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	k := 2
	println(divideArray(nums, k))
}
