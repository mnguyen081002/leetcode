package main

type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minstack) == 0 {
		this.minstack = append(this.minstack, val)
	} else if val <= this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	p := this.stack[len(this.stack)-1]
	if p == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minstack) == 0 {
		return 0
	}

	return this.minstack[len(this.minstack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
