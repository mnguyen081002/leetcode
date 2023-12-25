package main

func less(a1 []int, a2 []int) bool {
	return a1[0]*a1[0]+a1[1]*a1[1] < a2[0]*a2[0]+a2[1]*a2[1]
}

func partition(array [][]int, left, right int) int {
	x := left
	for i := left; i < right; i++ {
		if less(array[i], array[right]) {
			array[i], array[x] = array[x], array[i]
			x++
		}
	}

	array[x], array[right] = array[right], array[x]
	return x
}

func kClosest(points [][]int, k int) [][]int {
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
	inputs := [][]int{{1, 3}, {-2, 2}}
	kClosest(inputs, 1)
}
