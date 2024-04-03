package main

func numSubarraysWithSum(nums []int, goal int) int {
	l, r := 0, 0
	sum, ans := 0, 0
	for l <= r {
		sum += nums[r]
		r++
		for sum > goal && l < r {
			sum -= nums[l]
			l++
			if sum == goal {
				ans++
			}
		}
		if sum == goal {
			ans++
		}
	}
	return ans
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	goal := 2
	numSubarraysWithSum(nums, goal)
}
