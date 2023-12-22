package main

func containsDuplicate(nums []int) bool {
	set := map[int]uint8{}
	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = 0
	}
	return false
}
