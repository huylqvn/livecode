package livetest

import "container/heap"

// Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3

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
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)
	minHeap := &SetHeap{}
	heap.Init(minHeap)
	for i, v1 := range nums1 {
		heap.Push(minHeap, Set{sum: v1 + nums2[0], i: i, j: 0})
	}

	for minHeap.Len() > 0 && len(result) < k {
		set := heap.Pop(minHeap).(Set)
		result = append(result, []int{nums1[set.i], nums2[set.j]})
		if set.j+1 < len(nums2) {
			heap.Push(minHeap, Set{sum: nums1[set.i] + nums2[set.j+1], i: set.i, j: set.j + 1})
		}
	}

	return result
}

type Set struct {
	sum int
	i   int
	j   int
}

type SetHeap []Set

func (h SetHeap) Len() int           { return len(h) }
func (h SetHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h SetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h SetHeap) Empty() bool        { return len(h) == 0 }

func (h *SetHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

func (h *SetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
