package main

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for i := 2; i < n; i++ {
		l, r := 0, i-1
		for l < r {
			s := nums[l] + nums[r]
			if s <= nums[i] {
				l++
			} else {
				ans += r - l
				r--
			}
		}
	}
	return ans
}
