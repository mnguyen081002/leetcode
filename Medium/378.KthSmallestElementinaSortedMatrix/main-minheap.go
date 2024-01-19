package main

import "container/heap"

type MinHeap []Obj

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Obj))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Obj struct {
	val int
	r   int
	c   int
}

func kthSmallest1(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	h := &MinHeap{}
	for r := 0; r < min(m, k); r++ {
		heap.Push(h, Obj{matrix[r][0], r, 0})
	}
	for i := 0; i < k-1; i++ {
		obj := heap.Pop(h).(Obj)
		if obj.c+1 < n {
			heap.Push(h, Obj{matrix[obj.r][obj.c+1], obj.r, obj.c + 1})
		}
	}

	return heap.Pop(h).(Obj).val
}
