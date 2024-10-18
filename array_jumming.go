package livetest

import (
	"container/list"
)

type node struct {
	index int
	count int
}

func ArrayJumping(arr []int) int {
	target, index := 0, 0
	found := make(map[int]bool)

	// Find the largest number
	for x := 0; x < len(arr); x++ {
		found[x] = false
		if arr[x] > target {
			target = arr[x]
			index = x
		}
	}

	lIndex := index - target
	for lIndex < 0 {
		lIndex += len(arr)
	}
	rIndex := index + target
	for rIndex >= len(arr) {
		rIndex -= len(arr)
	}

	l := node{index: lIndex, count: 1}
	r := node{index: rIndex, count: 1}

	list := list.New()
	list.PushBack(l)
	list.PushBack(r)

	// BFS
	for list.Len() > 0 {
		cur := list.Front().Value.(node)
		list.Remove(list.Front())

		if cur.index == index {
			return cur.count
		}

		found[cur.index] = true

		// Calculate l and r indices
		lIndex = cur.index - arr[cur.index]
		for lIndex < 0 {
			lIndex += len(arr)
		}
		rIndex = cur.index + arr[cur.index]
		for rIndex >= len(arr) {
			rIndex -= len(arr)
		}

		l.index = lIndex
		l.count = cur.count + 1
		r.index = rIndex
		r.count = cur.count + 1

		if !found[l.index] {
			found[l.index] = true
			list.PushBack(l)

		}
		if !found[r.index] {
			found[r.index] = true
			list.PushBack(r)
		}
	}

	return -1
}
