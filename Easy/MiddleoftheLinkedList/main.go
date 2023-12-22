package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	l := lenList(head)

	if l == 1 {
		return head
	}
	for i := 0; i < l/2; i++ {
		head = head.Next
	}

	return head
}

func lenList(head *ListNode) int {
	len := 0
	for head != nil {
		head = head.Next
		len++
	}

	return len
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	head := &ListNode{}
	insertNode(head, values)
	x := middleNode(head)

	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
}

func insertNode(head *ListNode, values []int) {
	for i, v := range values {
		head.Val = v

		if i != len(values)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
}
