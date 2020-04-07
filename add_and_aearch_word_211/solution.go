package main

import "fmt"

type WordDictionary struct {
	isLeaf   bool
	children [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.add(word, 0)
}

func (this *WordDictionary) add(word string, step int) {
	if len(word) == step {
		this.isLeaf = true
		return
	}

	idx := word[step] - 'a'
	if this.children[idx] == nil {
		this.children[idx] = &WordDictionary{}
	}
	this.children[idx].add(word, step+1)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.doSearch(word, 0)
}

func (this *WordDictionary) doSearch(word string, step int) bool {
	if this == nil {
		return false
	} else if this.isLeaf == true && len(word) == step {
		return true
	} else if len(word) == step {
		return false
	}

	if word[step] == '.' {
		for i := 0; i < len(this.children); i++ {
			if this.children[i] != nil {
				if this.children[i].doSearch(word, step+1) {
					return true
				}

			}
		}
		return false
	}
	idx := word[step] - 'a'
	return this.children[idx].doSearch(word, step+1)
}

func main() {
	wd := Constructor()
	wd.AddWord("a")
	wd.AddWord("ab")
	fmt.Println(wd.Search(".a"))
	fmt.Println(wd.Search("ab"))
}
