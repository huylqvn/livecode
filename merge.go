package livetest

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 && n != 0 {
		nums1 = make([]int, n)
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if m == 0 || n == 0 {
		return
	}

	i, j := 0, 0
	res := []int{}
	for nums1[i] != 0 {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if j < n {
		for j < m {
			res = append(res, nums2[j])
			j++
		}
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = res[i]
	}
}
