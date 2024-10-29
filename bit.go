package livetest

import (
	"strconv"
)

func addBinary(a string, b string) string {
	// 11 + 1 = 100

	var res string
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	return res
}

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n &= n - 1
	}

	return n
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
