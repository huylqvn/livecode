package livetest

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				l++
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}
