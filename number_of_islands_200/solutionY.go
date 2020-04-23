package number_of_islands_200

func numIslands(grid [][]byte) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				counter++
				mark(grid, i, j)
			}
		}
	}
	return counter
}

func mark(grid [][]byte, row, col int) {
	if row < 0 || col < 0 ||
		row >= len(grid) || col >= len(grid[0]) {
		return
	}
	if grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'
	mark(grid, row+1, col)
	mark(grid, row-1, col)
	mark(grid, row, col+1)
	mark(grid, row, col-1)
}
