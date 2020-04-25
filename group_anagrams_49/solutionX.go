package group_anagrams_49

import "sort"

type Bytes []byte

func (bs Bytes) Len() int {
	return len(bs)
}

func (bs Bytes) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs Bytes) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func groupAnagrams(strs []string) [][]string {
	mapping := map[string][]string{}
	for _, str := range strs {
		bs := []byte(str)
		sort.Sort(Bytes(bs))
		key := string(bs)
		mapping[key] = append(mapping[key], str)
	}

	var result [][]string
	for _, val := range mapping {
		result = append(result, val)
	}
	return result
}
