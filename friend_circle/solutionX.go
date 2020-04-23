package main

func findCircleNum(m [][]int) int {
	parents := make([]int, len(m))
	for i := range parents {
		parents[i] = -1
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 1 && i != j {
				union(parents, i, j)
			}
		}
	}

	counter := 0
	for i := 0; i < len(parents); i++ {
		if parents[i] == -1 {
			counter++
		}
	}
	return counter
}

func find(parents []int, i int) int {
	idx := i
	for parents[idx] != -1 {
		idx = parents[idx]
	}
	root := idx

	for parents[i] != -1 {
		next := parents[i]
		parents[i] = root
		i = next
	}

	return root
}

func union(parents []int, i, j int) {
	pi := find(parents, i)
	pj := find(parents, j)
	if pi != pj {
		parents[pi] = pj
	}
}
