package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow_pointer, fast_pointer := head, head
	for fast_pointer != nil && fast_pointer.Next != nil {
		slow_pointer = slow_pointer.Next
		fast_pointer = fast_pointer.Next.Next
		if slow_pointer == fast_pointer {
			return true
		}
	}
	return false
}
