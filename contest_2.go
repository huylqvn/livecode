package livetest

func maxFactorScore(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	a := findLCM(nums)
	b := findGCD(nums)
	maxScore := int64(a) * int64(b)
	for i := 1; i < len(nums); i++ {
		newArr := removeElement(nums, i)
		if len(newArr) == 0 {
			continue
		}
		a := findLCM(newArr)
		b := findGCD(newArr)
		maxScore = max(maxScore, int64(a)*int64(b))
	}

	return maxScore
}

func removeElement(nums []int, i int) []int {
	res := []int{}
	for j := 0; j < len(nums); j++ {
		if i != j {
			res = append(res, nums[j])
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func findLCM(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func findGCD(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = gcd(result, arr[i])
	}
	return result
}

// z to ab
func lengthAfterTransformations(s string, t int) int {
	lenght := 0
	for i := 0; i < len(s); i++ {
		val := int(s[i] - 'a')
		distance := val + t - 26
		if distance < 0 {
			lenght++
		} else {
			loop := distance / 26
			p := 2
			for j := 1; j < loop+1; j++ {
				p *= 2
			}

			lenght += p
		}
	}

	return lenght
}
