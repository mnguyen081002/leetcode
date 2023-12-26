package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}
	if nums[0] > 0 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		h := len(nums) - 1
		for j < h {
			sum := nums[i] + nums[j] + nums[h]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[h]})
				j++
				h--

				for j < h && nums[j] == nums[j-1] {
					j++
				}

				for j < h && nums[h] == nums[h+1] {
					h--
				}
			} else if sum < 0 {
				j++
			} else {
				h--
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}
