package main

func colorTheArray(n int, queries [][]int) (ans []int) {
	nums, ans := make([]int, n), make([]int, len(queries))
	x := 0
	for i, q := range queries {
		index, color := q[0], q[1]
		if index > 0 && nums[index] > 0 && nums[index-1] == nums[index] {
			x--
		}
		if index < n-1 && nums[index] > 0 && nums[index+1] == nums[index] {
			x--
		}
		if index > 0 && nums[index-1] == color {
			x++
		}
		if index < n-1 && nums[index+1] == color {
			x++
		}
		ans[i] = x
		nums[index] = color
	}
	return
}
