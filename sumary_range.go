package livetest

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			i++
		}
		if start != nums[i] {
			res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d", start))
		}
	}
	return res

}
