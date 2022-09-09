package test828

func uniqueLetterString(s string) int {
	var mp = make(map[int32][]int)
	for idx, v := range s {
		if _, exist := mp[v]; !exist {
			mp[v] = make([]int, 0)
			mp[v] = append(mp[v], -1)
		}
		mp[v] = append(mp[v], idx)
	}
	var res = 0
	for _, v := range mp {
		v = append(v, len(s))
		for i := 1; i < len(v); i++ {
			res += (v[i] - v[i-1]) * (v[i+1] - v[i])
		}
	}
	return res
}
