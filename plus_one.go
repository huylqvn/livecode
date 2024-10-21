package livetest

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	last := digits[len(digits)-1]
	if last < 9 {
		digits[len(digits)-1]++
		return digits
	}
	return append(plusOne(digits[:len(digits)-1]), 0)
}

func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)/2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
