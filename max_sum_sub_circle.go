package livetest

func maxSubarraySumCircular(nums []int) int {
	maxSum := nums[0]
	minSum := nums[0]
	curMax := 0
	curMin := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// kadane algorithm
		curMax = max(curMax+nums[i], nums[i])
		maxSum = max(maxSum, curMax)

		// kadane algorithm
		curMin = min(curMin+nums[i], nums[i])
		minSum = min(minSum, curMin)
	}

	if maxSum < 0 {
		return maxSum
	}

	return max(maxSum, sum-minSum)
}
