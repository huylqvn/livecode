package livetest

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, h := 0, len(nums)-1

	first := -1
	last := -1

	for l <= h {
		m := (l + h) / 2
		if nums[m] == target {
			first = m
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	l, h = 0, len(nums)-1

	for l <= h {
		m := (l + h) / 2
		if nums[m] == target {
			last = m
			l = m + 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	if first == -1 || last == -1 {
		return []int{-1, -1}
	}

	return []int{first, last}
}
