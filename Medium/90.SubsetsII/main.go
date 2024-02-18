package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	sort.Ints(nums)
	backtracking(nums, []int{}, &res)
	return res
}

func backtracking(nums []int, cur []int, res *[][]int) {
	*res = append(*res, nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		*res = append(*res, append([]int{}, cur...))
		backtracking(nums[i+1:], cur, res)
		nums = nums[:len(nums)-1]
	}
}
