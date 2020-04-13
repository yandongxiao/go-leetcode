package palindromic_substrings_647

func countSubstrings(str string) int {
	counter := 0
	for i := 0; i < len(str); i++ {
		counter += do(str, i, i)
		counter += do(str, i, i+1)
	}
	return counter
}

func do(str string, left, right int) int {
	counter := 0
	for left >= 0 && right < len(str) && str[left] == str[right] {
		counter++
		left--
		right++
	}
	return counter
}
