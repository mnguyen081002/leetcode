package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, prev := head.Next, head
	for curr != nil {
		if curr.Val == prev.Val {
			curr = curr.Next
			prev.Next = curr
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return head
}
