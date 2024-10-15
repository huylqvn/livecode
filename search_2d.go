package livetest

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for _, m := range matrix {
		res := findTarget(m, target)
		if res {
			return true
		}
	}
	return false
}

func findTarget(nums []int, target int) bool {
	if target < nums[0] {
		return false
	}
	l, h := 0, len(nums)-1
	for l <= h {
		m := (l + h) / 2

		fmt.Println(l, h, m)
		if target == nums[m] {
			return true
		}
		if target >= nums[m] {
			l = m + 1
		} else {
			h = m
		}
	}
	return false
}
