package livetest

func validPath(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 {
		return true
	}
	mp := make(map[int][]int)
	for _, e := range edges {
		if v, ok := mp[e[0]]; !ok {
			mp[e[0]] = []int{e[1]}
		} else {
			v = append(v, e[1])
			mp[e[0]] = v
		}
	}

	if bfs(mp, -1, source, destination) {
		return true
	}

	if bfs(mp, -1, destination, source) {
		return true
	}

	return false
}

func bfs(mp map[int][]int, st, s, d int) bool {
	if st == s {
		return false
	}
	if st == -1 {
		st = s
	}
	if val, ok := mp[s]; ok {
		for _, v := range val {
			if v == d {
				return true
			}
		}
		for _, v := range val {
			res := bfs(mp, st, v, d)
			if res {
				return true
			}
		}
	}
	return false
}
