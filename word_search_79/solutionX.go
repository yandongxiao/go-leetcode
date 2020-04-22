package word_search_79

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		if exist(board, words[i]) {
			result = append(result, words[i])
		}
	}
	return result
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if helper(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, row, col int, word string) bool {
	if len(word) == 0 {
		return true
	} else if row >= len(board) || row < 0 {
		return false
	} else if col >= len(board[0]) || col < 0 {
		return false
	} else if board[row][col] == '_' || board[row][col] != word[0] {
		return false
	}

	prev := board[row][col]
	board[row][col] = '_'
	if helper(board, row+1, col, word[1:]) ||
		helper(board, row-1, col, word[1:]) ||
		helper(board, row, col+1, word[1:]) ||
		helper(board, row, col-1, word[1:]) {
		board[row][col] = prev
		return true
	}
	board[row][col] = prev

	return false
}
