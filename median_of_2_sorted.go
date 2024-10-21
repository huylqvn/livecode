package livetest

import "sort"

// Brute Force solution
func findMedianSortedArraysBF(nums1 []int, nums2 []int) float64 {
	res := append(nums1, nums2...)
	sort.Ints(res)
	if len(res)%2 == 0 {
		return float64(res[len(res)/2]+res[len(res)/2-1]) / 2
	}
	return float64(res[len(res)/2])
}

// binary search solution
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	imin, imax, halfLen := 0, m, (m+n+1)/2

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums2[j], nums1[i])
			}

			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
