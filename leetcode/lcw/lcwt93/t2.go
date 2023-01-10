package lcwt93

import (
	"sort"
)

func maxStarSum(vals []int, edges [][]int, k int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var res = vals[0]
	for _, v := range vals {
		if v > res {
			res = v
		}
	}

	for key, nodes := range g {
		var cur = []int{}
		for _, node := range nodes {
			cur = append(cur, vals[node])
		}
		sort.Ints(cur)
		cnt, total := 0, vals[key]
		for i := len(cur) - 1; i >= 0; i-- {
			if cnt >= k || cur[i] < 0 {
				break
			}
			total += cur[i]
			cnt++
		}
		if total > res {
			res = total
		}
	}
	return res
}
