package longest_palindromic_substring_5

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		s1 := expandPalindrome(s, i, i)
		s2 := expandPalindrome(s, i, i+1)
		max = Max(max, s1, s2)
	}
	return max
}

func expandPalindrome(s string, i, j int) string {
	if i < 0 || j+1 > len(s) || s[i] != s[j] {
		return ""
	}

	for i >= 0 && j <= len(s)-1 {
		if s[i] != s[j] {
			return s[i+1 : j]
		}
		i--
		j++
	}

	return s[i+1 : j]
}

func Max(str1, str2, str3 string) string {
	max := str1
	if len(max) < len(str2) {
		max = str2
	}
	if len(max) < len(str3) {
		max = str3
	}
	return max
}
