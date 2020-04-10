package longest_repeating_character_replacement_424

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	begin := 0
	val := s[begin]
	use := 0
	maxLen := 0
	diffIdx := -1
	for i := 1; i < len(s); i++ {
		if s[i] != val {
			if use < k {
				use++
				if diffIdx == -1 {
					diffIdx = i
				}
			} else {
				use = 0
				maxLen = max(maxLen, i-begin)
				if diffIdx != -1 {
					begin = diffIdx
					val = s[begin]
				} else {
					begin = i
					val = s[begin]
				}
				i = begin
				diffIdx = -1
			}
		}
	}

	added := min(k-use, begin)
	maxLen = max(maxLen, len(s)-begin+added)
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
