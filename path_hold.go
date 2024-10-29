package livetest

// l1 = "..XX.X."
// l2 = "X.X.X.."
func PathHold(s1, s2 string) int {
	n := len(s1)
	dp1, dp2 := 0, 0

	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			if s1[i] == 'X' {
				dp1++
				dp2++
			} else {
				dp1, dp2 = max(dp1, dp2), max(dp1, dp2)
			}
		} else {
			if s1[i] == 'X' {
				dp1, dp2 = max(dp1, dp2), dp2+1
			} else {
				dp2, dp1 = max(dp1, dp2), dp1+1
			}
		}
	}

	return max(dp1, dp2)
}
