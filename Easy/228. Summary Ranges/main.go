package main

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	ans := []string{}
	curr := string(nums[0] + '0')
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			ans = append(ans, curr+"->"+string(nums[i-1]+'0'))
			curr = string(nums[i] + '0')
		}
	}
	ans = append(ans, curr)
	return ans
}
