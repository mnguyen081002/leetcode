package main

// func removeDuplicates(nums []int) int {
// 	c, w := 1, 0

// 	for i := 1; i < len(nums)-1; i++ {
// 		left := c - 1
// 		right := i + 1
// 		if left >= 0 && right < len(nums) && nums[left] == nums[i] && nums[right] == nums[i] {
// 			nums[i], nums[w] = nums[w], nums[i]
// 			w++
// 		} else {
// 			nums[i], nums[c] = nums[c], nums[i]
// 			c++
// 			w++
// 		}
// 	}

// 	return c
// }

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	pos := 2
	prev2, prev := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		n := nums[i]
		if prev2 != n {
			nums[pos] = n
			pos++
		}
		prev2, prev = prev, n
	}
	return pos
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3}
	removeDuplicates(nums)
}
