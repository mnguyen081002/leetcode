package main

import (
	"container/heap"
)

type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	var maxHeap MaxHeap

	for _, v := range points {
		heap.Push(&maxHeap, v)
		if maxHeap.Len() > k {
			heap.Pop(&maxHeap)
		}
	}

	return maxHeap
}
