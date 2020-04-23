package main

func generateParenthesis(n int) []string {
	result := map[string]struct{}{}
	generate(s, n, n, result)
	output := make([]string, 0, len(result))
	for s := range result {
		output = append(output, s)
	}
	return output
}

func generate(s string, left, right int, result map[string]struct{}) {
	if left == 0 && right == 0 {
		result[s] = struct{}{}
		return
	}

	if left > 0 {
		generate(s+"(", left-1, right, result)
	}
	if right > left {
		generate(s+")", left, right-1, result)
	}
}
