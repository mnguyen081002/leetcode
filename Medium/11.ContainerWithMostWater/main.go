package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	maxwater := 0

	for left < right {
		maxwater = max(min(height[left], height[right])*(right-left), maxwater)
		if height[left] == height[right] {
			left++
			right--
		} else if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxwater
}
