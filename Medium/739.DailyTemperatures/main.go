package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	for index, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			distance := index - lastIndex
			result[lastIndex] = distance
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}

	return result
}
