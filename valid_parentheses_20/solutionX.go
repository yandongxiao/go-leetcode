package valid_parentheses_20

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
			continue
		}

		if len(stack) == 0 {
			return false
		}
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if (b == ')' && last != '(') ||
			(b == ']' && last != '[') ||
			(b == '}' && last != '{') {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
