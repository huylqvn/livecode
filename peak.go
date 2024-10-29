package livetest

// Input: nums = [1,2,3,1]
// Output: index =2

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	max := 0
	maxIndex := 0
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > max {
			max = nums[mid]
			maxIndex = mid
		}
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return maxIndex
}
