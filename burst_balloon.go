package livetest

import "sort"

func findMinArrowShots(points [][]int) int {
	ans := [][]int{}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	start := points[0][0]
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			ans = append(ans, []int{start, end})
			start = points[i][0]
			end = points[i][1]
		} else {
			if points[i][1] < end {
				end = points[i][1]
			}
		}
	}
	ans = append(ans, []int{start, end})
	return len(ans)
}
