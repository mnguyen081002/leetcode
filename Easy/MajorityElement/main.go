package main

import (
	"fmt"
	"math"
)

func majorityElement(nums []int) int {
	majorityEle, majorityFrequency := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == majorityEle {
			majorityFrequency++
		} else if majorityFrequency == 0 {
			majorityEle, majorityFrequency = nums[i], 1
		} else {
			majorityFrequency--
		}

		if majorityFrequency == int(math.Round(float64(len(nums))/2)) {
			return nums[i]
		}
	}

	return majorityEle
}

func main() {
	nums := []int{3, 3, 4}
	fmt.Println(majorityElement(nums))
}
