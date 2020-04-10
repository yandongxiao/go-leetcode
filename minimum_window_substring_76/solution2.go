package minimum_window_substring_76

import "container/heap"

type MinHeap []int

func (heap MinHeap) Less(i, j int) bool {
	return heap[i] < heap[j]
}

func (heap MinHeap) Len() int {
	return len(heap)
}

func (heap MinHeap) swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MinHeap) Push(x interface{}) {
	val := x.(int)
	*heap = append(*heap, val)
}

func (heap *MinHeap) Pop() interface{} {
	old := *heap
	last := len(old) - 1
	element := old[last]
	*heap = (*heap)[:last]
}

type MaxHeap []int

func (heap MaxHeap) Less(i, j int) bool {
	return heap[i] > heap[j]
}

func (heap MaxHeap) Len() int {
	return len(heap)
}

func (heap MaxHeap) swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MaxHeap) Push(x interface{}) {
	val := x.(int)
	*heap = append(*heap, val)
}

func (heap *MaxHeap) Pop() interface{} {
	old := *heap
	last := len(old) - 1
	element := old[last]
	*heap = (*heap)[:last]
}

func minWindow(s string, t string) string {
	set := map[byte]struct{}{}
	for i := 0; i < len(t); i++ {
		set[t[i]] = struct{}{}
	}

	bestMin := -1
	bestMax := -1
	minHeap := MinHeap{}
	maxHeap := MaxHeap{}
	heap.Init(&minHeap)
	heap.Init(&maxHeap)
	for i := 0; i < len(s); i++ {
		x := s[i]
		if _, ok := set[x]; ok {
			if old, ok := mapping[x]; ok {
				if min == -1 || min == old {

				}
			}
			mapping[x] = i
		} else {
			continue
		}

		min = Min(min, i)
		max = Max(min, i)
		if len(mapping) < len(t) {
			continue
		}
		if bestMin == -1 || bestMax == -1 || max-min < bestMax-bestMin {
			bestMin = min
			bestMax = max
		}
	}

	if bestMin == -1 || bestMax == -1 {
		return ""
	}
	return s[bestMin : bestMax+1]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
