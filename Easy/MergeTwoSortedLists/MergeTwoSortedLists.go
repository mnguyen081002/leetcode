/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// list1 = [1,2,4], list2 = [1,3,4]
// 1.list1 = 1, list1.Next = 2 -> list1 = 2, list1.Next = 4
// 2.list1 = 2, list2 = 4 -> list1 = 3, list2 = 4
// 3.list1 = 2, list2 = 3 -> list1 = 4, list2 = 3
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
		}
	}

	return list1
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	x := mergeTwoLists(list1, list2)

	for x != nil {
		println(x.Val)
		x = x.Next
	}
}
