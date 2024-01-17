package main

import "slices"

func permute(nums []int) [][]int {
	res := [][]int{}
	backtracking(nums, []int{}, &res)
	return res
}

func backtracking(nums []int, cur []int, res *[][]int) {
	if len(cur) == len(nums) {
		nw := make([]int, len(cur))
		copy(nw, cur)
		*res = append(*res, nw)
	} else {
		for i := 0; i < len(nums); i++ {
			if slices.Contains(cur, nums[i]) {
				continue
			}
			cur = append(cur, nums[i])
			backtracking(nums, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func main() {
	r := permute([]int{1, 2, 3})

	for _, v := range r {
		for _, vv := range v {
			print(vv, " ")
		}
		println()
	}
}
