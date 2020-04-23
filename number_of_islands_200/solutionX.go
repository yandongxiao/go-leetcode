package number_of_islands_200

type Point struct {
	row int
	col int
}

func numIslands(grid [][]byte) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				counter++
				mark(grid, Point{i, j})
			}
		}
	}
	return counter
}

func mark(grid [][]byte, p Point) {
	queue := []Point{p}
	for len(queue) > 0 {
		p = queue[0]
		queue = queue[1:]
		if p.row >= len(grid) ||
			p.row < 0 ||
			p.col >= len(grid[p.row]) ||
			p.col < 0 ||
			grid[p.row][p.col] == '0' {
			continue
		}
		grid[p.row][p.col] = '0'
		queue = append(queue, Point{p.row + 1, p.col})
		queue = append(queue, Point{p.row - 1, p.col})
		queue = append(queue, Point{p.row, p.col + 1})
		queue = append(queue, Point{p.row, p.col - 1})
	}
}
