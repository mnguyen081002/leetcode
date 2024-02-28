package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var m = map[*ListNode]struct{}{}
	if headA == nil || headB == nil {
		return nil
	}
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, f := m[headB]; f {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
