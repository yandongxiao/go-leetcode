package valid_parentheses_20

type Stack []byte

func (s *Stack) push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) pop() byte {
	old := *s
	last := len(old) - 1
	element := old[last]
	(*s) = old[:last]
	return element
}

func isValid(s string) bool {
	var stack Stack
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '{' || b == '(' || b == '[' {
			stack.push(b)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		p := stack.pop()
		if !(b == '}' && p == '{' ||
			b == ')' && p == '(' ||
			b == ']' && p == '[') {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
