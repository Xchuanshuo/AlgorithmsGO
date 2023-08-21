package s2

import (
	"math"
	"sort"
)

func minInterval(vs [][]int, queries []int) []int {
	sort.Slice(vs, func(i, j int) bool {
		if vs[i][1] != vs[j][1] {
			return vs[i][1] < vs[j][1]
		}
		return vs[i][1]-vs[i][0] > vs[j][1]-vs[j][0]
	})
	var ids = make([]int, len(queries))
	for i := range queries {
		ids[i] = i
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[ids[i]] < queries[ids[j]]
	})
	l, n := 0, len(vs)
	var res = make([]int, len(queries))
	for _, id := range ids {
		res[id] = math.MaxInt32
		var v = queries[id]
		for l < n && vs[l][1] < v {
			l++
		}
		if l == n || vs[l][0] > v {
			res[id] = -1
			continue
		}
		for ; l < n; l++ {
			res[id] = min(res[id], vs[l][1]-vs[l][0]+1)
			if !(l+1 < n && vs[l+1][0] <= v) {
				break
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
