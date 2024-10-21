package livetest

func longestConsecutive(nums []int) int {
	mp := make(map[int]int)

	for _, v := range nums {
		mp[v] = 1
	}

	max := 0
	for v := range mp {
		if mp[v-1] == 1 {
			continue
		}
		length := 1
		for mp[v+length] == 1 {
			length++
		}
		if length > max {
			max = length
		}
	}

	return max
}
