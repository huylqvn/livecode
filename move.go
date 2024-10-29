package livetest

func movesSolution(s string) int {
	moves := 0

	if len(s) < 2 {
		return len(s)
	}
	if s[0] != '>' {
		moves++
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '>' && i != len(s)-1 {
			continue
		}
		if s[i] == '<' && s[i-1] == '>' {
			continue
		}
		moves++
	}
	return moves
}
