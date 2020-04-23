package main

import "sort"

type timestamps [][]int

func (t timestamps) Len() int {
	return len(t)
}

func (t timestamps) Less(i, j int) bool {
	if t[i][0] < t[j][0] {
		return true
	}
	return false
}

func (t timestamps) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func earliestAcq(logs [][]int, N int) int {
	sort.Sort(timestamps(logs))

	parents := make([]int, N)
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}

	for i := 0; i < len(logs); i++ {
		person1 := logs[i][1]
		person2 := logs[i][2]
		if union(parents, person1, person2) {
			if numberOfParrents(parents) == 1 {
				return logs[i][0]
			}
		}
	}
	return -1
}

func numberOfParrents(parents []int) int {
	counter := 0
	for i := 0; i < len(parents); i++ {
		if parents[i] == -1 {
			counter++
		}
	}
	return counter
}

func find(parents []int, i int) int {
	for parents[i] != -1 {
		i = parents[i]
	}
	return i
}

func union(parents []int, i, j int) bool {
	pi := find(parents, i)
	pj := find(parents, j)
	if pi != pj {
		parents[pi] = pj
		return true
	}
	return false
}
