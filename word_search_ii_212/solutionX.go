package main

import "fmt"

func findWords(board [][]byte, words []string) []string {
	trie := &Trie{}
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}

	result := make(map[string]struct{})
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			helper(board, i, j, trie, result)
		}
	}

	data := make([]string, 0, len(result))
	for key := range result {
		data = append(data, key)
	}
	return data
}

func helper(board [][]byte, row, col int, trie *Trie, result map[string]struct{}) bool {
	if row >= len(board) || row < 0 {
		return false
	} else if col >= len(board[0]) || col < 0 {
		return false
	} else if board[row][col] == '_' {
		return false
	}

	x := board[row][col] - 'a'
	if trie.children[x] == nil {
		return false
	}
	if trie.children[x].word != "" {
		result[trie.children[x].word] = struct{}{}
	}

	prev := board[row][col]
	board[row][col] = '_'
	if helper(board, row+1, col, trie.children[x], result) ||
		helper(board, row-1, col, trie.children[x], result) ||
		helper(board, row, col+1, trie.children[x], result) ||
		helper(board, row, col-1, trie.children[x], result) {
		board[row][col] = prev
		return true
	}
	board[row][col] = prev
	return false
}

type Trie struct {
	word     string
	children [26]*Trie
}

func (t *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		x := word[i] - 'a'
		if t.children[x] == nil {
			t.children[x] = &Trie{}
		}
		t = t.children[x]
	}
	t.word = word
}

func main() {
	val := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words := []string{"ab", "ac"}
	fmt.Println(findWords(val, words))
}
