package word_search_79

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

func helper(board [][]byte, word string, step int, i, j int) bool {
	if step == len(word) {
		return true
	}

	if i < 0 || i >= len(board) {
		return false
	}

	if j < 0 || j >= len(board[0]) {
		return false
	}

	if word[step] != board[i][j] {
		return false
	}
	original := board[i][j]
	board[i][j] = '_'

	b := helper(board, word, step+1, i+1, j) ||
		helper(board, word, step+1, i-1, j) ||
		helper(board, word, step+1, i, j+1) ||
		helper(board, word, step+1, i, j-1)
	board[i][j] = original

	return b
}
