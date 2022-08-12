package test1282

func groupThePeople(groupSizes []int) [][]int {
	var mp = make(map[int][]int)
	var res = make([][]int, 0)
	for i := 0; i < len(groupSizes); i++ {
		var v = groupSizes[i]
		if _, exist := mp[v]; !exist {
			mp[v] = make([]int, 0)
		}
		mp[v] = append(mp[v], i)
		if len(mp[v]) == v {
			res = append(res, mp[v])
			delete(mp, v)
		}
	}
	return res
}
