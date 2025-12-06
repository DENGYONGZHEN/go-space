package stack

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

type MinStack struct {
	stack   []int
	min     []int
	pointer int
}

func Constructor() MinStack {
	return MinStack{
		stack:   []int{},
		min:     []int{},
		pointer: 0,
	}

}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.pointer == 0 {
		this.min = append(this.min, val)
	} else {
		minVal := this.GetMin()
		if val < minVal {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, minVal)
		}
	}
	this.pointer++

}

func (this *MinStack) Pop() {
	this.pointer--
	this.min = this.min[:this.pointer]
	this.stack = this.stack[:this.pointer]
}

func (this *MinStack) Top() int {
	return this.stack[this.pointer-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.pointer-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
