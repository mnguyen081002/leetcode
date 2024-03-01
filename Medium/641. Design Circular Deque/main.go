package main

type MyCircularDeque struct {
	Head      *LinkList
	Tail      *LinkList
	Length    int
	MaxLength int
}

type LinkList struct {
	Next *LinkList
	Val  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		MaxLength: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Length == this.MaxLength {
		return false
	}
	if this.Head == nil {
		this.Head = &LinkList{Val: value}
		this.Tail = this.Head
	} else {
		old := this.Head
		this.Head = &LinkList{Val: value, Next: old}
	}
	this.Length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Length == this.MaxLength {
		return false
	}
	if this.Head == nil {
		this.Head = &LinkList{Val: value}
		this.Tail = this.Head
	} else {
		this.Tail.Next = &LinkList{Val: value}
		this.Tail = this.Tail.Next
	}
	this.Length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.Length == 0 {
		return false
	}
	this.Head = this.Head.Next
	this.Length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.Length == 0 {
		return false
	}
	h := this.Head
	for h.Next != nil {
		if h.Next.Next == nil {
			h.Next = nil
			this.Tail = h
			break
		}
		h = h.Next
	}
	this.Length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.Length == 0 {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularDeque) GetRear() int {
	if this.Length == 0 {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Length == this.MaxLength
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
	// ["MyCircularDeque","insertFront","insertLast","getFront","insertLast","getFront","insertFront","getRear","getFront","getFront","deleteLast","getRear"]
	m := Constructor(5)
	m.InsertFront(7)
	m.InsertLast(0)
	m.GetFront()
	m.InsertLast(3)
	m.GetFront()
	m.InsertFront(9)
	m.GetRear()
	m.GetFront()
	m.GetFront()
	m.DeleteLast()
	m.GetRear()
}
