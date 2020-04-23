package main

import "fmt"

func generateParenthesis(n int) []string {
	result := map[string]struct{}{}
	generate(n, "", result)

	output := make([]string, 0, len(result))
	for s := range result {
		output = append(output, s)
	}
	return output
}

func generate(n int, s string, result map[string]struct{}) {
	if len(s) > 2*n {
		return
	} else if len(s) == 2*n {
		if valid(s) {
			result[s] = struct{}{}
			fmt.Println(s)
		}
		return
	}
	generate(n, s+"(", result)
	generate(n, s+")", result)
}

func valid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) <= 0 {
				return false
			}
			stack = stack[1:]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(generateParenthesis(3))
}
