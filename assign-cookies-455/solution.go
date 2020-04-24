package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	j := 0
	counter := 0
	for _, x := range g {
		for ; j < len(s); j++ {
			if s[j] < x {
				continue
			}
			counter++
			j++
			break
		}
		if j == len(s) {
			break
		}
	}
	return counter
}
