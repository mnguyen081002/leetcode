package main

import "sort"

func lengthOfLIS(nums []int) int {
	var piles [][]int
	for _, num := range nums {
		if len(piles) == 0 {
			piles = append(piles, []int{num})
			continue
		}
		j := sort.Search(len(piles), func(k int) bool {
			return piles[k][len(piles[k])-1] >= num
		})
		if j < len(piles) {
			piles[j] = append(piles[j], num)
		} else {
			piles = append(piles, []int{num})
		}
	}
	return len(piles)
}
