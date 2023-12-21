package main

type Stack struct {
	store []int
}

func (s *Stack) Push(x int) {
	s.store = append(s.store, x)
}

func (s *Stack) Pop() int {
	if len(s.store) == 0 {
		return -1
	}
	x := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return x
}

func (s *Stack) Peek() int {
	if len(s.store) == 0 {
		return -1
	}
	return s.store[len(s.store)-1]
}

func (s *Stack) Empty() bool {
	return len(s.store) == 0
}

func (s *Stack) Length() int {
	return len(s.store)
}

type MyQueue struct {
	stack1 Stack
	stack2 Stack
}

func Constructor() MyQueue {
	queue := MyQueue{
		stack1: Stack{},
		stack2: Stack{},
	}

	return queue
}

func (this *MyQueue) Push(x int) {
	l := this.stack2.Length()

	for i := 0; i < l; i++ {
		this.stack1.Push(this.stack2.Pop())
	}

	this.stack1.Push(x)
}

func (this *MyQueue) Pop() int {
	l := this.stack1.Length()

	for i := 0; i < l; i++ {
		pv := this.stack1.Pop()
		this.stack2.Push(pv)
	}

	v := this.stack2.Pop()

	for i := 0; i < l-1; i++ {
		this.stack1.Push(this.stack2.Pop())
	}

	return v
}

func (this *MyQueue) Peek() int {
	l := this.stack1.Length()
	for i := 0; i < l; i++ {
		pv := this.stack1.Pop()
		this.stack2.Push(pv)
	}
	v := this.stack2.Peek()
	for i := 0; i < l; i++ {
		x := this.stack2.Pop()
		this.stack1.Push(x)
	}
	return v
}

func (this *MyQueue) Empty() bool {
	return this.stack1.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	myQueue := Constructor()
	myQueue.Push(1) // queue is: [1]
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Peek()  // return 1
	myQueue.Pop()   // return 1, queue is [2]
	myQueue.Empty() // return false
}
