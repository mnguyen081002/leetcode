package main

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	sf := func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	}
	slices.SortFunc(intervals, sf)

	merge := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		if merge[1] < intervals[i][0] {
			result = append(result, merge)
			merge = intervals[i]
		} else if merge[0] > intervals[i][1] {
			result = append(result, intervals[i])
		} else {
			merge = []int{min(merge[0], intervals[i][0]), max(merge[1], intervals[i][1])}
		}
	}

	return append(result, merge)
}

func main() {
	nums := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	merge(nums)
}
