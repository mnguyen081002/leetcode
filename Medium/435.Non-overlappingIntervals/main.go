package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	remove := 0
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		lastIndex := len(merged) - 1
		if merged[lastIndex][1] > interval[0] {
			remove++
			continue
		}
		merged = append(merged, interval)
	}
	return len(intervals) - len(merged)
}

func main() {
	nums := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(nums))
}
