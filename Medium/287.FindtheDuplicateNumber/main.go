package main

import "fmt"

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	return 0
}

func main() {
	fmt.Println(10 % 10)
}
