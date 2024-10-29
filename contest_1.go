package livetest

// Input: word = "abbcccc"
// Output: 5
// Explanation:
// The possible strings are: "abbcccc", "abbccc", "abbcc", "abbc", and "abcccc"
func possibleStringCount(word string) int {
	count := 0
	for i := 0; i < len(word); {
		start := i
		for i < len(word) && word[i] == word[start] {
			i++
		}
		l := i - start

		if l > 1 {
			count += l - 1
		}
	}

	return count + 1
}

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	children := make([][]int, n)

	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	var dfsReparent func(node int)
	dfsReparent = func(node int) {
		for _, child := range children[node] {
			ancestor := parent[child]
			for ancestor != -1 && s[ancestor] != s[child] {
				ancestor = parent[ancestor]
			}
			if ancestor != -1 && ancestor != node {
				children[ancestor] = append(children[ancestor], child)
				children[node] = removeSlice(children[node], child)
				parent[child] = ancestor
			}

			parent[child] = node
			dfsReparent(child)
		}
	}

	var calculateSubtreeSizes func(node int) int
	calculateSubtreeSizes = func(node int) int {
		size := 1
		for _, child := range children[node] {
			size += calculateSubtreeSizes(child)
		}
		return size
	}

	dfsReparent(0)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = calculateSubtreeSizes(i)
	}

	return result
}

func removeSlice(slice []int, item int) []int {
	newSlice := []int{}
	for _, v := range slice {
		if v != item {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) // 0: ở lại, 1: di chuyển
		}
	}

	// Trường hợp cơ sở: ngày 1
	for i := 0; i < n; i++ {
		dp[i][0][0] = stayScore[0][i]
	}

	// Quy hoạch động
	for day := 1; day < k; day++ {
		for city := 0; city < n; city++ {
			// Ở lại
			dp[city][day][0] = stayScore[day][city] + max(dp[city][day-1][0], dp[city][day-1][1])

			// Di chuyển
			maxScore := 0
			for prevCity := 0; prevCity < n; prevCity++ {
				if prevCity != city {
					maxScore = max(maxScore, dp[prevCity][day-1][0]+travelScore[prevCity][city])
				}

				maxScore = max(maxScore, dp[city][day-1][1]+travelScore[city][prevCity])
			}
			dp[city][day][1] = maxScore

		}
	}

	// Tìm kết quả cuối cùng
	maxScore := 0
	for i := 0; i < n; i++ {
		maxScore = max(maxScore, max(dp[i][k-1][0], dp[i][k-1][1]))
	}
	return maxScore
}

func findOriginalTypedString(word string, k int) int {
	mod := 1000000007
	n := len(word)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// Base case: empty string
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		dp[i][0] = 1 // Empty string can be formed in one way
		for j := 1; j <= min(i, k); j++ {
			count := 0

			// Count the number of ways to form the prefix
			for p := 1; p <= min(j, 3); p++ {
				if i-p >= 0 && word[i-1] == word[i-p] {
					count = (count + dp[i-p][j-p]) % mod
				}
			}

			// Count the number of ways to form the string
			dp[i][j] = count
		}
	}

	return dp[n][k]
}
