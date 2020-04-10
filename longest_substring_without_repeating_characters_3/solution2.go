package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 0
	begin := 0
	mapping := map[byte]int{s[0]: 0}
	for i := 1; i < len(s); i++ {
		x := s[i]
		if k, ok := mapping[x]; ok {
			maxLen = max(maxLen, i-begin)
			for j := begin; j <= k; j++ {
				delete(mapping, s[j])
			}
			begin = k + 1
		}
		mapping[x] = i
	}

	return max(maxLen, len(s)-begin)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
