package test2246

/**
 * @Description https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/
 * idea: 树的直径（节点受限）
 **/

func longestPath(parent []int, s string) int {
	var g = make(map[int][]int)
	for i, p := range parent {
		g[p] = append(g[p], i)
	}
	var dfs func(fa, cur int) int
	var res = 1
	dfs = func(fa, cur int) int {
		var val = 0
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			var nxtV = dfs(cur, nxt)
			if s[nxt] != s[cur] {
				res = max(res, val+nxtV+1)
				val = max(val, nxtV)
			}
		}
		return val + 1
	}
	dfs(-1, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
