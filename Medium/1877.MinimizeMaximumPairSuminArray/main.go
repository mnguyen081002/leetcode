package main

import (
	"math"
	"slices"
)

func minPairSum(nums []int) int {
	slices.Sort(nums)
	res := math.MinInt
	for l, h := 0, len(nums)-1; l < h; l, h = l+1, h-1 {
		if s := nums[l] + nums[h]; s > res {
			res = s
		}
	}
	return res
}
