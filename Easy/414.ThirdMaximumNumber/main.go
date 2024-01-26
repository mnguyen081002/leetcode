package main

import "math"

func thirdMax(nums []int) int {
	mx1, mx2, mx3 := math.MinInt, math.MinInt, math.MinInt

	for _, v := range nums {
		if v > mx1 {
			mx1, mx2, mx3 = v, mx1, mx2
		} else if v > mx2 && v != mx1 {
			mx2, mx3 = v, mx2
		} else if v > mx3 && v != mx2 && v != mx1 {
			mx3 = v
		}
	}
	if mx3 != math.MinInt {
		return mx3
	}

	return mx1
}
