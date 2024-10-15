package livetest

func FindDuplicate(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := (low + high) / 2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
