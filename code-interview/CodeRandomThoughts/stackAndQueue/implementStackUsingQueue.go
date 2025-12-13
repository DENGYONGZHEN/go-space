package stackandqueue

// 225. Implement Stack using Queues

// Implement a last-in-first-out (LIFO) stack using only two queues.
// The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).
// Implement the MyStack class:

// void push(int x) Pushes element x to the top of the stack.
// int pop() Removes the element on the top of the stack and returns it.
// int top() Returns the element on the top of the stack.
// boolean empty() Returns true if the stack is empty, false otherwise.
// Notes:
// You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
// Depending on your language, the queue may not be supported natively.
// You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.

type MyStack struct {
	q []int
}

func StackConstructor() MyStack {
	return MyStack{q: []int{}}
}

// 关键在于考虑到在每次push时，把队列中的数依次放到新push的数的后面
func (s *MyStack) Push(x int) {
	s.q = append(s.q, x)
	// 把前面的元素转到队尾
	for i := 0; i < len(s.q)-1; i++ {
		s.q = append(s.q, s.q[0])
		s.q = s.q[1:]
	}
}

func (s *MyStack) Pop() int {
	top := s.q[0]
	s.q = s.q[1:]
	return top
}

func (s *MyStack) Top() int {
	return s.q[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}
