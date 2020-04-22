package main

type Trie struct {
	end      bool
	Children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		x := word[i] - 'a'
		if t.Children[x] == nil {
			t.Children[x] = &Trie{}
		}
		t = t.Children[x]
	}
	t.end = true
}

func (t *Trie) Search(word string) bool {
	if t = startsWith(t, word); t == nil {
		return false
	} else if !t.end {
		return false
	}
	return true
}

func (t *Trie) StartsWith(word string) bool {
	if startsWith(t, word) == nil {
		return false
	}
	return true
}

func startsWith(t *Trie, word string) *Trie {
	for i := 0; i < len(word); i++ {
		x := word[i] - 'a'
		if t.Children[x] == nil {
			return nil
		}
		t = t.Children[x]
	}
	return t
}
