package main

import (
	"slices"
	"strconv"
)

func largestNumber(nums []int) string {
	s := ""

	for _, v := range nums {
		s += strconv.Itoa(v)
	}

	ss := []byte(s)
	slices.SortFunc(ss, func(i, j byte) int {
		return int(j) - int(i)
	})

	return string(ss)
}
