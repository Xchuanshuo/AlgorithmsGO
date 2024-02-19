package lcw369

/**
 * @Description https://leetcode.cn/contest/weekly-contest-369/problems/maximum-points-after-collecting-coins-from-all-nodes/
 * idea: 树形dp
 **/

func maximumPoints(edges [][]int, coins []int, k int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var n = len(edges)
	var mem = make([][15]int, n+1)
	for i := 0; i < len(mem); i++ {
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}
	var dfs func(fa, cur int, cnt int) int
	dfs = func(fa, cur int, cnt int) int {
		if cnt >= 14 {
			return 0
		}
		if mem[cur][cnt] != -1 {
			return mem[cur][cnt]
		}
		// dp1当前不/2; dp2当前/2；分别计算两种情况的最大值
		var dp1 = coins[cur]>>cnt - k
		var dp2 = coins[cur] >> (cnt + 1)
		for _, nxt := range g[cur] {
			if fa == nxt {
				continue
			}
			dp1 += dfs(cur, nxt, cnt)
			dp2 += dfs(cur, nxt, cnt+1)
		}
		mem[cur][cnt] = max(dp1, dp2)
		return max(dp1, dp2)
	}
	return dfs(-1, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
