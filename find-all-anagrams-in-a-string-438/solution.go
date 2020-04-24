package main

func findAnagrams(s string, p string) []int {
	pmap := make([]int, 26)
	for _, b := range p {
		b = b - 'a'
		pmap[b]++
	}

	smap := make([]int, 26)
	for i := 0; i < len(p) && i < len(s); i++ {
		b := s[i] - 'a'
		smap[b]++
	}

	var result []int
	if equal(smap, pmap) {
		result = append(result, 0)
	}

	idx := 0
	for i := len(p); i < len(s); i++ {
		smap[s[i]-'a']++
		smap[s[idx]-'a']--
		idx++
		if equal(smap, pmap) {
			result = append(result, idx)
		}
	}
	return result
}

func equal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
