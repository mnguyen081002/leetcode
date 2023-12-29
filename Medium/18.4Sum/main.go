package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, 0, 4, target)
}

// key point = kSum(nums, i+1, k-1, target-nums[i])
func kSum(nums []int, start, k, target int) [][]int {
	res := [][]int{}
	// không tìm thấy || 4 min của array lớn hơn target || 4 max của array nhỏ hơn target
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
	low, high := start, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum < target || (low > start && nums[low] == nums[low-1]) {
			low++
		} else if sum > target || (high < len(nums)-1 && nums[high] == nums[high+1]) {
			high--
		} else {
			res = append(res, []int{nums[low], nums[high]})
			low++
			high--
		}
	}
	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fourSum(nums, 0)
}
