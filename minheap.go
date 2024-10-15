package livetest

import (
	"container/heap"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func initHeap() *minHeap {
	minHeap := &minHeap{}
	heap.Init(minHeap)
	return minHeap
}

func HandleMinHeap(a []int) int {
	h := initHeap()
	for _, v := range a {
		heap.Push(h, v)
	}
	return heap.Pop(h).(int)
}
