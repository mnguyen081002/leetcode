package main

import "container/heap"

type Fraction struct {
	a, b int
}
type MaxHeap []Fraction

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return float64(h[i].a)/float64(h[i].b) > float64(h[j].a)/float64(h[j].b)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Fraction))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	mh := MaxHeap{}
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(&mh, Fraction{arr[i], arr[j]})
			if len(mh) > k {
				heap.Pop(&mh)
			}
		}
	}
	ans := make([]int, 2)
	x := heap.Pop(&mh).(Fraction)
	ans[0], ans[1] = x.a, x.b
	return ans
}

func main() {
	arr := []int{1, 7}
	k := 1
	println(kthSmallestPrimeFraction(arr, k))
}
