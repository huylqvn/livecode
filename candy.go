package livetest

func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	for i := 0; i < len(ratings)-1; i++ {
		if candies[i] == 0 {
			candies[i] = 1
		}
		if ratings[i] < ratings[i+1] {
			candies[i+1] = candies[i] + 1
		}
	}

	for i := len(ratings) - 1; i >= 1; i-- {
		if candies[i] == 0 {
			candies[i] = 1
		}
		if ratings[i] < ratings[i-1] {
			candies[i-1] = max(candies[i]+1, candies[i-1])
		}
	}

	sum := 0
	for i := 0; i < len(candies); i++ {
		sum += candies[i]
	}
	return sum
}
