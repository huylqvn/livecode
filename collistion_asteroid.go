package livetest

func AsteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, n := range asteroids {
		if len(stack) > 0 {
			v := n
			check := false
			for {
				if v == 0 {
					break
				}
				if len(stack) == 0 && v != 0 {
					stack = append(stack, v)
					break
				}
				prev := v
				sval := stack[len(stack)-1]
				v, check = collistion(sval, prev)
				if v == sval {
					break
				}
				if !check {
					stack = append(stack, prev)
					break
				}

				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, n)
		}
	}
	return stack
}

func collistion(a, b int) (int, bool) {
	if a > 0 && b > 0 {
		return 0, false
	}
	if a < 0 && b < 0 {
		return 0, false
	}
	if a < 0 && b > 0 {
		return 0, false
	}
	x, y := a, b

	if b < 0 && a > 0 {
		y = -b
	}

	if x > y {
		return a, true
	}
	if x == y {
		return 0, true
	}
	return b, true
}
