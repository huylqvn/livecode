package livetest

import (
	"sort"
)

func FixHole(a []int) int {
	if len(a) < 3 {
		return 1
	}
	sort.Ints(a)
	res := 1<<31 - 1
	for i := 1; i < len(a); i++ {
		res = min(res, maxSub(a[:i], a[i:]))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSub(a, b []int) int {
	if len(a) == 1 {
		return b[len(b)-1] - b[0]
	}
	if len(b) == 1 {
		return a[len(a)-1] - a[0]
	}
	x := a[len(a)-1] - a[0]
	y := b[len(b)-1] - b[0]
	if x > y {
		return x
	}
	return y
}
