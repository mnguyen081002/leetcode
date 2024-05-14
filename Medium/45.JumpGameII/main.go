package main

func jump(nums []int) int {
	var window int
	var jumps int
	var furthest int

	for i := 0; i < len(nums)-1; i++ {
		jump := nums[i] + i
		if jump > furthest {
			furthest = jump
		}
		if i == window {
			jumps++
			window = furthest
		}
	}
	return jumps
}

func main() {
	jump([]int{2, 3, 1, 1, 4})
}
