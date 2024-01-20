package main

func findKthLargest(nums []int, k int) int {
	min_val, max_val := nums[0], nums[0]
	for _, n := range nums {
		min_val = min(min_val, n)
		max_val = max(max_val, n)
	}

	b := make([]int, max_val-min_val+1)
	for _, n := range nums {
		b[n-min_val]++
	}

	cntr, id := 0, len(b)-1
	for cntr < k {
		for b[id] == 0 {
			id--
		}
		b[id]--
		cntr++
	}
	return id + min_val
}

func main() {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
}
