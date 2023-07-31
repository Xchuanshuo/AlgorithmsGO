package test2328

import "sort"

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var M = int(1e9) + 7

func countPaths(g [][]int) int {
	m, n := len(g), len(g[0])
	var mp = make(map[int][][]int)
	var a = make([]int, 0)
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if len(mp[g[i][j]]) == 0 {
				a = append(a, g[i][j])
			}
			mp[g[i][j]] = append(mp[g[i][j]], []int{i, j})
		}
		dp[i] = make([]int, n)
	}
	sort.Ints(a)
	var res = 0
	for _, k := range a {
		for _, v := range mp[k] {
			x, y := v[0], v[1]
			dp[x][y] = 1
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || g[nx][ny] >= g[x][y] {
					continue
				}
				dp[x][y] = (dp[x][y] + dp[nx][ny]) % M
			}
			res = (res + dp[x][y]) % M
		}
	}
	return res
}
