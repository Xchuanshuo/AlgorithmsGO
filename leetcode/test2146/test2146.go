package test2146

import "sort"

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func highestRankedKItems(g [][]int, price []int, s []int, k int) [][]int {
	m, n := len(g), len(g[0])
	var vis = make([]bool, m*n)
	var dis = make([]int, m*n)
	var ta = make([][]int, 0)
	var q = [][]int{s}
	vis[s[0]*n+s[1]] = true
	var level = 0
	for len(q) > 0 {
		var sz = len(q)
		for i := 0; i < sz; i++ {
			var cur = q[0]
			q = q[1:]
			if g[cur[0]][cur[1]] >= price[0] && g[cur[0]][cur[1]] <= price[1] {
				dis[cur[0]*n+cur[1]] = level
				ta = append(ta, []int{cur[0], cur[1]})
			}
			for _, d := range dirs {
				nx, ny := cur[0]+d[0], cur[1]+d[1]
				var p = nx*n + ny
				if nx < 0 || nx >= m || ny < 0 || ny >= n || vis[p] || g[nx][ny] == 0 {
					continue
				}
				vis[p] = true
				q = append(q, []int{nx, ny})
			}
		}
		level++
	}
	sort.Slice(ta, func(i, j int) bool {
		x1, y1 := ta[i][0], ta[i][1]
		x2, y2 := ta[j][0], ta[j][1]
		d1, d2 := dis[x1*n+y1], dis[x2*n+y2]
		if d1 != d2 {
			return d1 < d2
		}
		if g[x1][y1] != g[x2][y2] {
			return g[x1][y1] < g[x2][y2]
		}
		if x1 != x2 {
			return x1 < x2
		}
		return y1 < y2
	})
	return ta[0:min(k, len(ta))]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
