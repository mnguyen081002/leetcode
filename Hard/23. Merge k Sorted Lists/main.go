package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (m MinHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}
func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}
func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MinHeap) Len() int {
	return len(m)
}
func mergeKLists(lists []*ListNode) *ListNode {
	h := MinHeap{}
	for _, l := range lists {
		for l != nil {
			curr := l
			l = l.Next
			curr.Next = nil
			h = append(h, curr)
		}
	}
	heap.Init(&h)
	l := &ListNode{}
	c := l
	for len(h) > 0 {
		l.Next = heap.Pop(&h).(*ListNode)
		l = l.Next
	}
	return c.Next
}
