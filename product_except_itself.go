package livetest

func productExceptSelf(nums []int) []int {
	length := len(nums)

	res := make([]int, length)

	prefix := 1
	for i := 0; i < length; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := length - 1; i > -1; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
