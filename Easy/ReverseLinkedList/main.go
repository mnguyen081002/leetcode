package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil // 1

	for {
		next := head.Next // 2
		head.Next = prev  // 2 -> 1

		prev = head
		if next == nil {
			break
		}
		head = next // 1 = 2
	}

	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	h := reverseList(head)

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
