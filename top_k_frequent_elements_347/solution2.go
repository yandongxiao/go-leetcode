package top_k_frequent_elements_347

import "container/heap"

type Item struct {
	key int
	val int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val > pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	last := len(old) - 1
	item := old[last]
	*pq = old[:last]
	return item
}

func topKFrequent(nums []int, k int) []int {
	mapping := make(map[int]int)
	for _, x := range nums {
		mapping[x]++
	}

	pq := make(PriorityQueue, k)
	heap.Init(&pq)
	for k, v := range mapping {
		heap.Push(&pq, &Item{
			key: k,
			val: v,
		})
	}

	var result []int
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.key)
	}
	return result
}
