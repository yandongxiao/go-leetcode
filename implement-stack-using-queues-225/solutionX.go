package main

package main

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q queue) peek() int {
	return q[len(q)-1]
}

func (q queue) empty() bool {
	return len(q) == 0
}

func (q queue) size() int {
	return len(q)
}

func Constructor() MyStack {
	return MyStack{}
}

type MyStack struct {
	q      queue
	backup queue
}

func (s *MyStack) Push(n int) {
	s.q.push(n)
}

func (s *MyStack) Pop() int {
	for s.q.size() > 1 {
		s.backup.push(s.q.pop())
	}
	result := s.q.pop()
	s.q, s.backup = s.backup, s.q
	s.backup = s.backup[:]
    return result
}

func (s *MyStack) Top() int {
	for s.q.size() > 1 {
		s.backup.push(s.q.pop())
	}
	result := s.q.peek()
	s.backup.push(s.q.pop())
	s.q, s.backup = s.backup, s.q
	s.backup = s.backup[:]
	return result
}

func (s *MyStack) Empty() bool {
	return s.q.empty()
}

