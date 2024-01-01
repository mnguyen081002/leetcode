package main

func canSeePersonsCount(heights []int) []int {
	stack := []int{}
	result := make([]int, len(heights))

	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] > stack[len(stack)-1] {
			result[i]++
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i]++
		}
		stack = append(stack, heights[i])
	}

	return result
}
