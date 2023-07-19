package test1733

func minimumTeachings(n int, lgs [][]int, friends [][]int) int {
	var m = len(lgs)
	var t = make([]map[int]bool, m+1)
	for i, lg := range lgs {
		t[i+1] = make(map[int]bool)
		for _, l := range lg {
			t[i+1][l] = true
		}
	}
	var newFriends = make([][]int, 0)
	for _, f := range friends {
		var fu = t[f[0]]
		var valid = false
		for k := range t[f[1]] {
			if fu[k] {
				valid = true
				break
			}
		}
		if !valid {
			newFriends = append(newFriends, f)
		}
	}
	var res = m
	// 枚举每门语言和没有共同语言的好友
	for i := 1; i <= n; i++ {
		var set = make(map[int]bool)
		for _, f := range newFriends {
			fu, fv := t[f[0]], t[f[1]]
			if !fu[i] {
				set[f[0]] = true
			}
			if !fv[i] {
				set[f[1]] = true
			}
		}
		res = min(res, len(set))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
