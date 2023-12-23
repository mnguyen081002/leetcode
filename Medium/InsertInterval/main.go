package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int

	for i, v := range intervals {
		if newInterval[1] < v[0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if newInterval[0] > v[1] {
			result = append(result, v)
		} else {
			newInterval = []int{
				min(newInterval[0], v[0]),
				max(newInterval[1], v[1]),
			}
		}
	}

	return append(result, newInterval)
}

func main() {
	intervals := [][]int{{1, 5}}
	newInterval := []int{0, 3}
	n := insert(intervals, newInterval)

	fmt.Println(n)
}
