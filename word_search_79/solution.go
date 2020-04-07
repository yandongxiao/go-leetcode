package word_search_79

// Time: O(N * 4^(L (len of word)))
// Space: O(L) recursion stack at most L
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if helper(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, word string, step int, i int, j int) bool {
	if step == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return false
	}

	if board[i][j] != word[step] {
		return false
	}

	original := board[i][j]
	board[i][j] = '_'

	if helper(board, word, step+1, i+1, j) ||
		helper(board, word, step+1, i-1, j) ||
		helper(board, word, step+1, i, j+1) ||
		helper(board, word, step+1, i, j-1) {
		return true
	}

	board[i][j] = original
	return false
}
