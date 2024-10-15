package livetest

func Permute(nums []int) [][]int {
	res := [][]int{}

	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}

	for i, n := range nums {
		x := n
		temp := remove(nums, i)

		pp := Permute(temp)
		for _, p := range pp {
			res = append(res, append([]int{x}, p...))
		}
	}

	return res
}

func remove(nums []int, i int) []int {
	res := []int{}
	for j := 0; j < len(nums); j++ {
		if i != j {
			res = append(res, nums[j])
		}
	}
	return res
}
