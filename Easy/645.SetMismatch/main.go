package main

func findErrorNums(nums []int) []int {
	a := make([]int, len(nums))
	d := 0
	for _, v := range nums {
		if a[v-1] != 0 {
			d = v
		} else {
			a[v-1] = v
		}
	}

	for i, v := range a {
		if v == 0 {
			return []int{d, i + 1}
		}
	}
	return []int{}
}

func main() {
	findErrorNums([]int{1, 2, 2, 4})
}
