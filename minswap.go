package livetest

func minSwaps(nums []int) int {
	onesCount := 0
	for _, v := range nums {
		if v == 1 {
			onesCount++
		}
	}

	maxOnesInWindow, currentOnes := 0, 0

	for i := 0; i < onesCount; i++ {
		currentOnes += nums[i]
	}
	maxOnesInWindow = currentOnes

	for i := onesCount; i < len(nums); i++ {
		currentOnes += nums[i] - nums[i-onesCount]
		maxOnesInWindow = max(maxOnesInWindow, currentOnes)
	}

	return onesCount - maxOnesInWindow
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
