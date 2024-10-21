package livetest

func longestValidParentheses(s string) int {
	stack := []int{-1}
	res := 0

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				size := i - stack[len(stack)-1]
				res = max(res, size)
			}
		}

	}

	return res

}
