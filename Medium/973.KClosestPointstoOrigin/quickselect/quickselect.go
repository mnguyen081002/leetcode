package main

import "fmt"

func partition(array []int, left, right int) int {
	x := left
	for i := left; i < right; i++ {
		if array[i] < array[right] {
			array[i], array[x] = array[x], array[i]
			x++
		}
	}

	array[x], array[right] = array[right], array[x]
	return x
}

func kClosest(points []int, k int) []int {
	var pivot int
	left := 0
	right := len(points) - 1
	for left < right {
		pivot = partition(points, left, right)
		if pivot-left == k-1 {
			break
		}
		if pivot > k {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return points[:k]
}

func main() {
	inputs := []int{4, 3, 9, -2, 8, 10}
	fmt.Println(kClosest(inputs, 3))
}
