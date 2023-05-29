package test1079

import "sort"

func numTilePossibilities(tiles string) int {
	var bs = []byte(tiles)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	var n = len(bs)
	var visited = make([]bool, n)
	var res = 0
	var dfs func()
	dfs = func() {
		res++
		for i := 0; i < n; i++ {
			if visited[i] || (i > 0 && bs[i-1] == bs[i] && !visited[i-1]) {
				continue
			}
			visited[i] = true
			dfs()
			visited[i] = false
		}
	}
	dfs()
	return res - 1
}
