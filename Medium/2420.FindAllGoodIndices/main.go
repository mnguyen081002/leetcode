package main

func goodIndices(nums []int, k int) []int {
	res := []int{}
	for i := k; i < len(nums)-k; i++ {
		if isNonInc(nums, i, k) && isNonDec(nums, i, k) {
			res = append(res, i)
		}
	}
	return res
}

func isNonInc(nums []int, i, k int) bool {
	for x := i - 1; x > i-k; x-- {
		if nums[x] > nums[x-1] {
			return false
		}
	}
	return true
}

func isNonDec(nums []int, i, k int) bool {
	for x := i + 1; x < i+k; x++ {
		if nums[x] > nums[x+1] {
			return false
		}
	}
	return true
}

func main() {
	goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2)
}
