package livetest

func dfs(initial int, n int, arr *[]int) {
	if initial > n {
		return
	} else {
		*arr = append(*arr, initial)
		for i := 0; i < 10; i++ {
			newnum := initial*10 + i
			if newnum <= n {
				dfs(newnum, n, arr)
			} else {
				return
			}
		}
	}
}

func LexicalOrder(n int) []int {
	var arr []int
	for i := 1; i < 10; i++ {
		dfs(i, n, &arr)
	}
	return arr
}
