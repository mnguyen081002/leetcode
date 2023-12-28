package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, 0, 4, target)
}

func kSum(nums []int, start, k, target int) [][]int {
	res := [][]int{}
	if start == len(nums) || nums[start]*k > target || target > nums[len(nums)-1]*k {
		return res
	}
	if k == 2 {
		return twoSum(nums, start, target)
	}
	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, set := range kSum(nums, i+1, k-1, target-nums[i]) {
				res = append(res, append([]int{nums[i]}, set...))
			}
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	res := [][]int{}
	lo, hi := start, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target || (lo > start && nums[lo] == nums[lo-1]) {
			lo++
		} else if sum > target || (hi < len(nums)-1 && nums[hi] == nums[hi+1]) {
			hi--
		} else {
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
		}
	}
	return res
}
