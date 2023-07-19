package test2747

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	var ids = make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[i]] < queries[ids[j]]
	})
	var res = make([]int, len(queries))
	var set = make(map[int]int)
	l, r := 0, 0
	for _, id := range ids {
		lv, rv := queries[id]-x, queries[id]
		for r < len(logs) && logs[r][1] <= rv {
			set[logs[r][0]]++
			r++
		}
		for l < len(logs) && logs[l][1] < lv {
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
