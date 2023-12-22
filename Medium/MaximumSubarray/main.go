package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	var maxSub = nums[0]
	var curSum = 0
	for _, num := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += num
		maxSub = max(maxSub, curSum)
	}

	return maxSub
}
