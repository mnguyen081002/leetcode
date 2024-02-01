package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	println(removeNthFromEnd(head, 1))

	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
