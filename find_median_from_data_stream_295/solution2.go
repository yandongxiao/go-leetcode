package find_median_from_data_stream_295

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (pq MaxHeap) Len() int {
	return len(pq)
}

func (pq MaxHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MaxHeap) Pop() interface{} {
	old := *pq
	last := len(old) - 1
	item := old[last]
	*pq = old[:last]
	return item
}

type MinHeap []int

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	last := len(old) - 1
	item := old[last]
	*pq = old[:last]
	return item
}

type MedianFinder struct {
	smallerHalf *MaxHeap
	largerHalf  *MinHeap
}

func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{
		smallerHalf: maxHeap,
		largerHalf:  minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.smallerHalf, num)
	heap.Push(this.largerHalf, heap.Pop(this.smallerHalf))
	if this.largerHalf.Len() > this.smallerHalf.Len() {
		heap.Push(this.smallerHalf, heap.Pop(this.largerHalf))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	fmt.Println(*this.smallerHalf, *this.largerHalf)
	// NOTE：只能访问下标为0的元素，其它下标的元素是没有排序的
	if this.largerHalf.Len() == this.smallerHalf.Len() {
		return float64((*this.largerHalf)[0]+(*this.smallerHalf)[0]) / 2
	}
	return float64((*this.smallerHalf)[0])
}
