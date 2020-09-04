package main

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	size := len(*s)
	result := (*s)[size-1]
	*s = (*s)[:size-1]
	return result

}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func (s *stack) peek() int {
	size := len(*s)
	return (*s)[size-1]
}

type MyQueue struct {
	input  stack
	output stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(n int) {
	q.input.push(n)
}

func (q *MyQueue) Peek() int {
	q.load()
	return q.output.peek()
}

func (q *MyQueue) Pop() int {
	q.load()
	return q.output.pop()
}

func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}

func (q *MyQueue) load() {
	if q.output.empty() {
		for !q.input.empty() {
			q.output.push(q.input.pop())
		}
	}
}
