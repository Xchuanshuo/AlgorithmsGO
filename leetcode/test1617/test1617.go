package test1617

/**
 * @Description https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/
 * idea: 枚举 + 求所有子树的最大直径
 **/

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	var g = make([][]int, n+1)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var dfs func(cur int) int
	var set = make(map[int]bool)
	var d = 0
	dfs = func(cur int) int {
		var curV = 0
		delete(set, cur)
		for _, nxt := range g[cur] {
			if !set[nxt] {
				continue
			}
			var nxtV = dfs(nxt)
			d = max(d, curV+nxtV)
			curV = max(curV, nxtV)
		}
		return curV + 1
	}
	var res = make([]int, n-1)
	for i := 1; i < 1<<n; i++ {
		set = make(map[int]bool)
		var s = 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				set[j+1] = true
				s = j + 1
			}
		}
		d = 0
		dfs(s)
		if len(set) == 0 && d > 0 {
			res[d-1]++
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
