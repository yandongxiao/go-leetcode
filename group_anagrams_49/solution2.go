package group_anagrams_49

import (
	"sort"
	"sync"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type ChanData struct {
	str string
	i   int
}

func sortStrings(strs []string) []string {
	var wg sync.WaitGroup
	wg.Add(len(strs))

	outputChan := make(chan *ChanData, len(strs))
	for i := 0; i < len(strs); i++ {
		go func(i int) {
			defer wg.Done()
			str := sortString(strs[i])
			outputChan <- &ChanData{
				str: str,
				i:   i,
			}
		}(i)
	}
	wg.Wait()
	close(outputChan)

	result := make([]string, len(strs))
	for output := range outputChan {
		result[output.i] = output.str
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	keys := sortStrings(strs)

	mapping := map[string][]string{}
	for i, str := range strs {
		key := keys[i]
		mapping[key] = append(mapping[key], str)
	}

	output := [][]string{}
	for _, val := range mapping {
		output = append(output, val)
	}
	return output
}
