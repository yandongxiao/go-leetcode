package main

type Trie struct {
	endOfWord bool
	children  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i <= len(word); i++ {
		if i == len(word) {
			node.endOfWord = true
			return
		}
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i <= len(word); i++ {
		if node == nil {
			return false
		}
		if node.endOfWord == true && len(word) == i {
			return true
		}
		if len(word) == i {
			return false
		}
		idx := word[i] - 'a'
		node = node.children[idx]
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.doStartsWith(prefix, 0)
}

func (this *Trie) doStartsWith(prefix string, step int) bool {
	if this == nil {
		return false
	} else if len(prefix) == step {
		return true
	}

	idx := prefix[step] - 'a'
	return this.children[idx].doStartsWith(prefix, step+1)
}
