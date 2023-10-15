package test2065

import "sort"

/**
 * @Description https://leetcode.cn/problems/maximum-path-quality-of-a-graph/
 * idea: bfs
 **/

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	var g = make(map[int][][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], []int{e[1], e[2]})
		g[e[1]] = append(g[e[1]], []int{e[0], e[2]})
	}
	var q = make([]T, 0)
	q = append(q, T{0, 0, []int{0}})
	var res = 0
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		if cur.v == 0 {
			sort.Ints(cur.nodes)
			var val = 0
			for i, v := range cur.nodes {
				if i == 0 || v != cur.nodes[i-1] {
					val += values[v]
				}
			}
			res = max(res, val)
		}
		for _, nxt := range g[cur.v] {
			if cur.times+nxt[1] > maxTime {
				continue
			}
			var dst = make([]int, len(cur.nodes))
			copy(dst, cur.nodes)
			q = append(q, T{cur.times + nxt[1], nxt[0], append(dst, nxt[0])})
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type T struct {
	times int
	v     int
	nodes []int
}
