package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{[]int{}}
	var cur []int

	sort.Ints(nums)
	doSubsetsWithDup(nums, cur, &res)
	return res
}

func doSubsetsWithDup(nums []int, cur []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		newRes := make([]int, len(cur))
		copy(newRes, cur)
		*res = append(*res, newRes)

		doSubsetsWithDup(nums[i+1:], cur, res)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	subsetsWithDup([]int{1, 2, 2})
}
