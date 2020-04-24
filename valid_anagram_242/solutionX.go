package valid_anagram_242

func isAnagram(s string, t string) bool {
	mapping := make(map[rune]int)
	for _, x := range s {
		mapping[x]++
	}

	for _, x := range t {
		mapping[x]--
	}

	for _, v := range mapping {
		if v != 0 {
			return false
		}
	}
	return true
}
