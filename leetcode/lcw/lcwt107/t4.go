package lcwt107

import "sort"

func countServers(n int, logs [][]int, x int, qs []int) []int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	var ids = make([]int, 0)
	for i := range qs {
		ids = append(ids, i)
	}
	sort.Slice(ids, func(i, j int) bool {
		return qs[ids[i]] < qs[ids[j]]
	})
	var set = make(map[int]int)
	l, r := 0, 0
	var res = make([]int, len(qs))
	for _, id := range ids {
		lv, rv := qs[id]-x, qs[id]
		for r < len(logs) && logs[r][1] <= rv {
			set[logs[r][0]]++
			r++
		}
		for l <= r && l < len(logs) && logs[l][1] < lv {
			set[logs[l][0]]--
			if set[logs[l][0]] == 0 {
				delete(set, logs[l][0])
			}
			l++
		}
		res[id] = n - len(set)
	}
	return res
}
