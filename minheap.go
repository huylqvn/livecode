package livetest

import (
	"container/heap"
)

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
