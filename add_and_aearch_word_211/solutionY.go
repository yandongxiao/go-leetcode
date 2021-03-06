package main

import "fmt"

type WordDictionary struct {
	end      bool
	children [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (wd *WordDictionary) AddWord(word string) {
	for i := 0; i < len(word); i++ {
		x := word[i] - 'a'
		if wd.children[x] == nil {
			wd.children[x] = &WordDictionary{}
		}
		wd = wd.children[x]
	}
	wd.end = true
}

func (wd *WordDictionary) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < len(wd.children); j++ {
				if wd.children[j] != nil {
					if wd.children[j].Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		}
		x := word[i] - 'a'
		if wd.children[x] == nil {
			return false
		}
		wd = wd.children[x]
	}
	if !wd.end {
		return false
	}
	return true
}

func main() {
	wd := Constructor()
	wd.AddWord("bad")
	fmt.Println(wd.Search("b.."))
}
