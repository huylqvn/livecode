package livetest

func MatchingCharacters(str string) int {
	mp := make(map[rune][]int)
	for i, char := range str {
		mp[char] = append(mp[char], i)
	}

	count := 0
	for _, mc := range mp {
		// only check if there are more than 1
		if len(mc) >= 2 {
			for i := 0; i < len(mc)-1; i++ {
				check := make(map[rune]bool)
				for j := mc[i] + 1; j < mc[len(mc)-1]; j++ {
					check[rune(str[j])] = true
				}
				if len(check) > count {
					count = len(check)
				}
			}
		}
	}

	return count
}
