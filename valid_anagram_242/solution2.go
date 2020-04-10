package valid_anagram_242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	book := map[rune]int
	for i, x := range s {
		book[x]++
	}

	for i, x := range t {
		book[x]--
	}

	for _, val := range book {
		if val != 0 {
			return false
		}
	}
	return true
}
