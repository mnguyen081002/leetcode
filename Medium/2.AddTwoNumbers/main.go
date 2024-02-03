package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d := 0
	r := &ListNode{}
	c := r
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			r.Next = &ListNode{Val: (l1.Val + l2.Val + d) % 10}
			d = (l1.Val + l2.Val + d) / 10
			r, l1, l2 = r.Next, l1.Next, l2.Next
		} else if l1 == nil {
			r.Next = &ListNode{Val: (d + l2.Val) % 10}
			d = (d + l2.Val) / 10
			r, l2 = r.Next, l2.Next
		} else {
			r.Next = &ListNode{Val: (d + l1.Val) % 10}
			d = (d + l1.Val) / 10
			r, l1 = r.Next, l1.Next
		}
	}
	if d != 0 {
		r.Next = &ListNode{Val: d}
	}

	return c.Next
}

func main() {
	l1 := &ListNode{Val: 3, Next: &ListNode{Val: 7}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 2}}

	for l := addTwoNumbers(l1, l2); l != nil; l = l.Next {
		println(l.Val)
	}
}
