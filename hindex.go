package livetest

func hIndex(citations []int) int {
	counter := make([]int, len(citations)+1)

	for _, c := range citations {
		if c > len(citations) {
			counter[len(citations)]++
		} else {
			counter[c]++
		}
	}

	sum := 0
	for i := len(citations); i >= 0; i-- {
		sum += counter[i]
		if sum >= i {
			return i
		}
	}
	return 0
}
