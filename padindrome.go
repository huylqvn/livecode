package livetest

func PalindromeSolution(s string, k int) string {
	n := len(s)
	if n < 1 {
		return s
	}
	if n < 2 {
		if s[0] == '?' {
			return "a"
		}
		return s
	}
	res := make([]byte, n)
	for i := 0; i < n/2; i++ {
		if s[i] != '?' && s[n-i-1] != '?' {
			if s[i] != s[n-i-1] {
				k--
				if k < 0 {
					return "NO"
				}
			}
			res[i] = s[i]
			res[n-i-1] = s[i]
		} else if s[i] == '?' && s[n-i-1] != '?' {
			res[i] = s[n-i-1]
			res[n-i-1] = s[n-i-1]
		} else if s[i] != '?' && s[n-i-1] == '?' {
			res[n-i-1] = s[i]
			res[i] = s[i]
		} else {
			res[i] = 'a'
			res[n-i-1] = 'a'
		}
	}

	if n%2 != 0 {
		if s[n/2] == '?' {
			res[n/2] = 'a'
		} else {
			res[n/2] = s[n/2]
		}
	}

	return string(res)
}
