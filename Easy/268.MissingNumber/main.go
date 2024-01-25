package main

func missingNumber(nums []int) int {
	l := len(nums)
	sum := (1 + l) * l / 2
	for _, num := range nums {
		sum -= num
	}

	return sum
}
