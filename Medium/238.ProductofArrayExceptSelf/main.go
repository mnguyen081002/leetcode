package main

func productExceptSelf(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	result[0] = 1
	for i := 1; i < size; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := 1
	for i := size - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	productExceptSelf(nums)
}
